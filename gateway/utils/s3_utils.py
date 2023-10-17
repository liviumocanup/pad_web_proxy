import boto3
import botocore
from boto3.exceptions import Boto3Error
from fastapi import HTTPException

from config.config import config

session = boto3.Session(
    aws_access_key_id=config.AWS_ACCESS_KEY,
    aws_secret_access_key=config.AWS_SECRET_KEY
)

s3 = session.resource('s3')


def create_unique_filename(user_id: str, file_name: str) -> str:
    return f"{user_id}_{file_name}"


def upload_file(file, unique_filename: str):
    if file_exists(config.BUCKET_NAME, unique_filename):
        raise HTTPException(status_code=400, detail="File with that name already exists.")

    try:
        # Upload the file to S3
        s3.Bucket(config.BUCKET_NAME).upload_fileobj(
            file.file,
            unique_filename,
            ExtraArgs={'ContentDisposition': f'attachment; filename="{file.filename}"'}
        )
    except Boto3Error as e:
        raise HTTPException(status_code=500, detail=str(e))


def delete_file(unique_filename: str):
    s3.Object(config.BUCKET_NAME, unique_filename).delete()


def file_exists(bucket_name, file_key):
    try:
        s3.Object(bucket_name, file_key).load()
    except botocore.exceptions.ClientError as e:
        if e.response['Error']['Code'] == "404":
            return False
        else:
            raise
    return True


def download_file(s3_url: str, destination_path: str):
    # Parsing the bucket name and object key from the URL
    bucket_name = s3_url.split('/')[2].split('.')[0]
    file_key = "/".join(s3_url.split('/')[3:])

    try:
        s3.Bucket(bucket_name).download_file(file_key, destination_path)
    except Boto3Error as e:
        raise HTTPException(status_code=500, detail=str(e))
