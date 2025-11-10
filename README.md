A simple command-line quiz game written in Go.

## 📘 Description
This project reads quiz questions from a CSV file, randomizes their order, and asks them one by one.  
Each line in `quiz.csv` contains a question and its answer, separated by a comma.


## 🚀 How to Run
```bash
go build main.go
./main
```

## Shell arguments
After building the go files
```bash
./main --help
```
This will show all the shell parameters supported:
```bash
  -csv string
        A CSV File in the format of 'questions,answer' (default "quiz.csv")
  -limit int
        Time limit for the quiz in seconds (default 30)
  -n int
        Max Number of questions in Quiz (default 20)
```

Different ways to run the code
```bash
go run main.go
go run main.go -csv="quiz.csv" -n=3 -limit=10
./main -csv="quiz.csv" -n=3 -limit=10
```

🧩 Features

- Reads questions and answers from a CSV file
- Randomizes question order
- Tracks score and shows incorrect answers
- Once the timer expires, quiz will terminate showing the final score.
