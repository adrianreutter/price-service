# Price Service

Calculates brutto price for a specific quantity.

## Endpoints

### POST /v1/price

Takes:

```
{
  "quantity": integer,
  "nettoPrice": float,
  "taxes": float
}
```

Gives:

```
{
  "quantity": integer,
  "nettoPrice": float,
  "taxes": float
  "calculatedBruttoPrice": float
}
```
