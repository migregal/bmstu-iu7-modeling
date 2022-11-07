import random
import time

random.seed(time.time())

import matplotlib.pyplot as plt
import scipy.integrate as integrate
import numpy as np
import pandas as pd
import networkx as nx


def find_steady_ps(matrix: np.ndarray) -> np.ndarray:
    b = np.array([0] * (len(matrix) - 1) + [1])

    matrix_to_solve = matrix.copy().transpose()
    matrix_to_solve -= np.diag(matrix.sum(axis=1))
    matrix_to_solve[-1] = np.ones(len(matrix_to_solve))

    return np.linalg.solve(matrix_to_solve, b)


def find_steady_ts(
    ts: np.ndarray, ps: np.ndarray, steady_ps: np.ndarray, eps: float = 1e-6
) -> np.ndarray:
    steady_ts = np.zeros(len(steady_ps))

    for i, row in enumerate(ps):
        for p, t in reversed(list(zip(row, ts))):
            if abs(steady_ps[i] - p) > eps:
                steady_ts[i] = t
                break

    return steady_ts


def solve_ode(matrix: np.ndarray, start_probs: np.ndarray, tn: int, tmax: int):
    matrix_to_solve = matrix.copy().transpose()
    matrix_to_solve -= np.diag(matrix.sum(axis=1))

    ts = np.linspace(0, tmax, tn)

    ps = integrate.odeint(
        lambda w, _: matrix_to_solve @ w, start_probs, ts, atol=1.0e-8, rtol=1.0e-6
    ).transpose()
    return ts, ps


def draw_graph(matrix) -> None:
    G = nx.from_numpy_matrix(matrix, create_using=nx.DiGraph)
    layout = nx.circular_layout(G)

    nx.draw(G, layout, node_size=600)
    nx.draw_networkx_labels(
        G,
        pos=layout,
        labels={i: f"$S_{i+1}$" for i in range(len(matrix))},
        font_color="yellow",
        font_size=12,
    )
    nx.draw_networkx_edge_labels(
        G, pos=layout, edge_labels=nx.get_edge_attributes(G, "weight"), label_pos=0.2
    )
    plt.show()


def draw_plot(ts: np.ndarray, ps: np.ndarray) -> None:
    for i, p in enumerate(ps):
        plt.plot(ts, p, label=f"$P_{{{i}}}$")
    plt.legend()
    plt.show()


def main():
    tn = 10**5
    tmax = 10

    m = np.array(
        [
            [0, 0, 1, 0],
            [4, 0, 3, 0],
            [0, 3, 0, 4],
            [0, 1, 1, 0],
        ]
    )

    draw_graph(m)

    start_probs = m[0].copy()

    ts, ps = solve_ode(m, start_probs, tn, tmax)

    draw_plot(ts, ps)

    steady_ps = find_steady_ps(m)
    steady_ts = find_steady_ts(ts, ps, steady_ps)

    df = pd.DataFrame.from_dict(
        {
            "Вероятностная константа": steady_ps,
            "Время достижения вероятностной константы": steady_ts,
        },
        orient="index",
        columns=[f"$S_{{{i+1}}}$" for i in range(len(m))],
    )

    df.style.format("{:.3f}")
    print(df)


if __name__ == "__main__":
    main()
