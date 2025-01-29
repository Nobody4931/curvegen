# Polynomial Curve Generator

A simple tool used to generate a polynomial function that passes through an arbitrary set of points

## How does it work?

Just a bit of simple algebra:

Let's define $F_n(x_i) = y_i$ for `i <= n`

This means `F_1(x) = y_1`

We can define the rest as `F_[n+1](x) = F_n(x) + G(x)`,
where `G(x_[n+1]) = y_[n+1] - F_n(x_[n+1])`
and `G(x_i) = 0` for `i <= n`

Therefore: `G(x) = Pi(i->n)(x - x_i) * (y_[n+1] - F_n(x_[n+1]) / Pi(i->n)(x_[n+1] - x_i)`
