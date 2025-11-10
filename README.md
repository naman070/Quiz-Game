A simple command-line quiz game written in Go.

## 📘 Description
This project reads quiz questions from a CSV file, randomizes their order, and asks them one by one.  
Each line in `quiz.csv` contains a question and its answer, separated by a comma.


## 🚀 How to Run
```bash
go run main.go
```

You can also limit the number of questions:
```bash
go run main.go 5
```

🧩 Features

- Reads questions and answers from a CSV file
- Randomizes question order
- Tracks score and shows incorrect answers
