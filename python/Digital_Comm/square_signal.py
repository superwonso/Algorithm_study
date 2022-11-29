from scipy import signal
import matplotlib.pyplot as plot
import numpy as np

A=1; g=3

# Define the rect function
def rect(x):
    y=x
    for i in range(len(x)):
        if np.abs(x[i]) < 0.5:
            y[i] = 1
        elif np.abs(x[i]) == 0.5:
            y[i] = 0.5
        else:
            y[i] = 0
    return y

# Sampling rate 60 hz / second
t = np.linspace(0, 50, 60, endpoint=True)

# Plot the square wave signal
plot.plot(t, A*rect(t/g))

# Give a title for the square wave plot
plot.title('Sqaure wave')

# Give x axis label for the square wave plot
plot.xlabel('T')

# Give y axis label for the square wave plot
plot.ylabel('Amplitude')
plot.grid(True, which='both')

# Provide x axis and line color
plot.axhline(y=0, color='k')

# Set the max and min values for y axis
plot.ylim(-2, 2)

# Display the square wave drawn
plot.show()