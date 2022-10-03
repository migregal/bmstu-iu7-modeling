import matplotlib.pyplot as plt
import numpy as np

data = np.genfromtxt("result.csv", delimiter=",", names=[
                     'X', '1', '2', '3', '4', 'e1', 'e2', 'rk'])

plt.plot(data['X'], data['1'], label='Picard, 1st approx')
plt.plot(data['X'], data['2'], label='Picard, 2nd approx')
plt.plot(data['X'], data['3'], label='Picard, 3rd approx')
plt.plot(data['X'], data['4'], label='Picard, 4th approx')

plt.legend()
plt.grid()
plt.gcf().set_size_inches(10, 4)
plt.savefig('pikar.png')

plt.clf()

plt.plot(data['X'], data['e1'], label='Euler')
plt.plot(data['X'], data['e2'], label='Euler, implicit')

plt.legend()
plt.grid()
plt.gcf().set_size_inches(10, 4)
plt.savefig('euler.png')

plt.clf()

plt.plot(data['X'], data['rk'], label='Runge-Kutt')

plt.legend()
plt.grid()
plt.gcf().set_size_inches(10, 4)
plt.savefig('rk.png')
