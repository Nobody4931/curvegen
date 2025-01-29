# Polynomial Curve Generator

A simple tool used to generate a polynomial function that passes through an arbitrary set of points

## How does it work?

Just a bit of simple algebra:

Let's define $F_n(x_i) = y_i$ for all $i \in [1, n]$, where:
- $(x_i, y_i)$ is a point in $\mathbb{R}^2$
- $F_k(x)$ is the polynomial function passing through all points $\{(x_i, y_i) : i \in [1, k]\}$

Therefore, $F_1(x) = y_1$.

Using this, we can use recursion to define the rest:

$F_{k + 1}(x) = F_k(x) + g(x)$, where:
- $g(x_{k + 1}) = y_{k + 1} - F_k(x_{k + 1})$ (in order to offset the error from $F_k(x)$)
- $g(x_i) = 0$ for all $i \in [1, k]$ (in order to not introduce any more error)

Taking these constraints into consideration, $g(x)$ can be defined as follows:

$g(x) = (\Pi_{i = 1}^{k} (x - x_i))(\frac{y_{k + 1} - F_n(x_{k + 1})}{\Pi_{i = 1}^{k} x_{k + 1} - x_i})$

## Okay that's cool and all but why do you need this?

Just for experimentation purposes. But hey, check out this funny little program I made in Java using this polynomial generator that prints out "hello":

```java
public class Main {
	public static void main(String[] args) {
		double a = Double.longBitsToDouble(0x409DB822E6B4883EL);
		double b = Double.longBitsToDouble(0xC0882EF8192C761DL);
		double c = Double.longBitsToDouble(0x405976E3A6223F14L);
		double d = Double.longBitsToDouble(0xC01375783F18F53AL);
		double e = Double.longBitsToDouble(0x3FB110944E7E4E42L);
		double f = Double.longBitsToDouble(0xBF2D05AFF955340AL);
		int g = 69, h,i;
		while (g > 0) System.out.print((char)((0x69 % 10*2) * ((1 << 3) + 2) + (g = (int)(g*b+a + (h=g*g)*(g*d+c)+e*(i=h*h) + f*i*g)) % (1 << 4) - 4));
	}
}
```

Concept inspired by some YouTube video I saw about funny ways to print out hello world.
