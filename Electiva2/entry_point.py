import joblib
import os
def model_fn(model_path):
    return joblib.load(os.path.join(model_path, 'model.joblib'))
