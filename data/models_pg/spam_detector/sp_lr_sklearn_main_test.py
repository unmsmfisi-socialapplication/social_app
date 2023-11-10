import unittest
from unittest.mock import patch
import sys
from execution.sp_lr_sklearn_main import run_model_training

sys.path.append('C:/Users/USUARIO/Documents/GitHub/social_app/data')

class TestMainScript(unittest.TestCase):

    @patch('schedule.every')
    @patch('schedule.run_pending')
    @patch('time.sleep')
    def test_run_model_training(self, mock_sleep, mock_run_pending, mock_every):
        # Llamar a la función principal y verificar que las funciones de schedule y time se llamen
        run_model_training()
        mock_every.assert_called_once()
        mock_run_pending.assert_called_once()
        mock_sleep.assert_called_once()

if __name__ == '__main__':
    unittest.main()
