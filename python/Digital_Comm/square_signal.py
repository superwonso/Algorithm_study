from scipy import signal
import matplotlib.pyplot as plot
import numpy as np

A=1; g=3

# Sampling rate 1000 hz / second
t = np.linspace(-20, 20, 1000, endpoint=True)

# Plot the square wave signal
plot.plot(t, A*signal.square(t/g))

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
plot.ylim(-1.5, 1.5)

# Display the square wave drawn
plot.show()