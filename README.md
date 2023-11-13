# exercise-go-gorm

16. Answer: tidak bisa seperti itu, karena jadinya 2 where tersebut akan menjadi AND, sehungga tidak sesuai ekspektasi yang inginnya dipisah. query yang ditunjukan bertujuan untuk dapat mengambil pet yang agresive sekaligus tidak agresive

20. Answer: Query pertama lebih tahan dari sql injection karena ada query berparameter yaitu yang menggunakan '?', sedangkan query kedua tidak menggunakan parameter '?' dan langsung passing variabel, sehingga rawan adanya sql injection