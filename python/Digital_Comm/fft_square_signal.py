from scipy import fft
from scipy import signal 
import matplotlib.pyplot as plot
import numpy as np

A=1; g=1

# Sampling rate 1000 hz / second
t = np.linspace(0, 20, 1000, endpoint=True)

# Calculate the FFT of the square wave signal
y_tmp = signal.square(t/g)
fft_y = np.fft.fft(y_tmp)

# Plot the fft square wave signal
plot.plot(t, fft_y)

# Give x axis label for the square wave plot
plot.xlabel('Frequency')

# Give y axis label for the square wave plot
plot.ylabel('τ')
plot.grid(True, which='both')

# Provide x axis and line color
plot.axhline(y=0, color='k')

# Set the max and min values for y axis
plot.ylim(-2, 2)

# Display the square wave drawn
plot.show()