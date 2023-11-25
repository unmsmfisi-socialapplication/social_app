import time
import schedule
from training.sp_lr_sklearn_training import run_model_training

if __name__ == "__main__":
    interval_seconds = 24 * 60 * 60  # 24 horas

    schedule.every(interval_seconds).seconds.do(run_model_training)

    run_model_training()

    while True:
        schedule.run_pending()
        time.sleep(1)
