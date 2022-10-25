import matplotlib.pyplot as plt
from scipy.stats import uniform, poisson
import numpy as np

print("Plot params")
x_from, x_to = float(input("Min X: ")), float(input("Max X: "))
steps = int(input("Steps: "))
x = np.linspace(x_from, x_to, steps)

print("Distribution params")
a, b = float(input("A: ")), float(input("B: "))

fig, axis = plt.subplots(1, 2, figsize=(12,7))
axis[0].set_title("F of Uniform Distribution")
axis[0].plot(x, uniform.cdf(x, a, b-a))
axis[1].set_title("f of Uniform Distribution")
axis[1].plot(x, uniform.pdf(x, a, b-a))
plt.show()

print("Plot params")
x_from, x_to = float(input("Min X: ")), float(input("Max X: "))
steps = int(input("Steps: "))
x = np.linspace(x_from, x_to, steps)
x_pmf = np.arange(int(x_from), np.ceil(x_to))

lmbd = float(input("λ (must be > 0): "))

fig, axis = plt.subplots(1, 2, figsize=(12,7))
axis[0].set_title("F of Poisson Distribution")
axis[0].vlines(x, 0, poisson.cdf(x, lmbd), 
               linestyles='-', lw=1)
axis[1].set_title("f of Poisson Distribution")
axis[1].vlines(x_pmf, 0, poisson.pmf(x_pmf, lmbd), 
               linestyles='-', lw=1)
plt.show()
