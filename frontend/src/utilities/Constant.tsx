export const nameRegex = /^[a-zA-ZÀ-ÿ\s'-]{1,150}$/;

export const emailRegex = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/;

export const passwordRegex =
  /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
