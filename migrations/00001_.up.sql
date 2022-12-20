CREATE TABLE IF NOT EXISTS currency
(
   currencyfrom text,
   currencyto text,
   well float default 0.0,
   updated_at timestamptz
)
