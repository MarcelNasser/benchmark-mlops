"""
Linear regression:
"""
import os
import unittest
from random import randint
import mlflow
from sklearn.linear_model import LinearRegression


class TestMlflowLinerReg(unittest.TestCase):


    def setUp(self) -> None:
        # server URL
        mlflow.set_tracking_uri(os.getenv('MLFLOW_URL'))
        # spin logger
        mlflow.autolog()
        # experiment
        mlflow.set_experiment(os.getenv('MLFLOW_EXPERIMENT'))

    def test_training(self):
        TRAIN_SET_LIMIT = 1000
        TRAIN_SET_COUNT = 100

        TRAIN_INPUT = list()
        TRAIN_OUTPUT = list()
        for i in range(TRAIN_SET_COUNT):
            a = randint(0, TRAIN_SET_LIMIT)
            b = randint(0, TRAIN_SET_LIMIT)
            c = randint(0, TRAIN_SET_LIMIT)
            op = a + (2 * b) + (3 * c)
            TRAIN_INPUT.append([a, b, c])
            TRAIN_OUTPUT.append(op)

        predictor = LinearRegression(n_jobs=-1)
        predictor.fit(X=TRAIN_INPUT, y=TRAIN_OUTPUT)

        X_TEST = [[10, 20, 30]]
        outcome = predictor.predict(X=X_TEST)
        coefficients = predictor.coef_

if __name__ == '__main__':
    unittest.main()
