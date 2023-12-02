CREATE TABLE Model_Code (
    ID INT PRIMARY KEY,
    Model NVARCHAR(100),
    NameTable NVARCHAR(100)
);


INSERT INTO Model_Code (ID, Model, NameTable)
VALUES (1, 'Post', 'bi_post_transformada'),
(2, 'Spam', 'bi_spam_transformada'),
(3, 'Sentimental', 'bi_sentimental_transformada');
