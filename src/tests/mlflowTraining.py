"""
Test Training and Metric Logging
"""
import os
import unittest
from random import random, randint

import mlflow
from sklearn.model_selection import train_test_split
from sklearn.datasets import load_diabetes
from sklearn.ensemble import RandomForestRegressor


class TestMlFlow(unittest.TestCase):

    def setUp(self) -> None:
        # server URL
        mlflow.set_tracking_uri(os.getenv('MLFLOW_URL'))
        # spin logger
        mlflow.autolog()
        # experiment
        mlflow.set_experiment(os.getenv('MLFLOW_EXPERIMENT'))

    def test_metrics_logging(self):
        mlflow.log_param("config_value", randint(0, 100))
        # Log a dictionary of parameters
        mlflow.log_params({"param1": randint(0, 100), "param2": randint(0, 100)})
        # Log a metric; metrics can be updated throughout the run
        mlflow.log_metric("accuracy", random() / 2.0)
        mlflow.log_metric("accuracy", random() + 0.1)
        mlflow.log_metric("accuracy", random() + 0.2)
        assert mlflow.get_artifact_uri()

    def test_training(self):
        db = load_diabetes()
        x_train, x_test, y_train, y_test = train_test_split(db.data, db.target)
        # Create and train models.
        rf = RandomForestRegressor(n_estimators=100, max_depth=6, max_features=3)
        rf.fit(x_train, y_train)
        # Use the model to make predictions on the test dataset.
        predictions = rf.predict(x_test)
        assert len(predictions) > 0


if __name__ == '__main__':
    unittest.main()
