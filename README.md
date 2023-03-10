# loan-schedule-API

REST API Calculates the borrower's schedule.

## Run the app

    go run ./main.go

## Run the tests

    go test ./...

# REST API

Runs on port 3000

## Calculates the borrower schedule

| Parameter                     | Type    | Description                                              |
| ----------------------------- | ------- | -------------------------------------------------------- |
| `capital emprunté`            | `int`   | **Required** borrowed capital                            |
| `taux d'intérêt annuel`       | `float` | **Required** annual interest rate                        |
| `nombre d'échéance`           | `int`   | **Required** number of terms before the end of the loan  |

```http
POST /
```

- Response

There are as many elements in the answer as there are terms

```json
[
    {
        "montant de l échéance"         : int,
        "capital restant dû"            : int,
        "montant des intérêts"          : int,
        "montant du capital"            : int,
        "capital restant à rembourser"  : int
    },
    ...
]
```