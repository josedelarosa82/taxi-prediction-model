payload = '1, 2, 1\n 2, 3, 5\n'

import boto3

endpoint = 'sagemaker-scikit-learn-2022-12-06-01-34-50-163'

client = boto3.client('sagemaker-runtime','us-east-1')

response = client.invoke_endpoint(
    EndpointName=endpoint,
    Body=payload,
    ContentType='text/csv'
)

prediction = response['Body'].read()
print(prediction)