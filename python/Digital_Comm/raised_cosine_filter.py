import matplotlib.pyplot as plot
import numpy as np

r = 0 # roll-off factor , r = 0, 0.5, 1
W_0 = 0.5; f_delta = r*W_0 # normalized frequency

# Define the raised cosine function
# P(f)= 1 (abs(f)<=W_0 - f_delta) , (1+cos(pi*abs(f)-f_1/2f_delta))/2 (W_0-f_delta<abs(f)<=W_0+f_delta) , 0 (abs(f)>W_0+f_delta)
def raised_cosine(f, W_0, f_delta):
    y = np.zeros(len(f))
    if r == 1: 
        return np.sinc(f)
    for i in range(len(f)):
        if np.abs(f[i]) <= W_0 - f_delta:
            y[i] = 1
        elif (W_0-f_delta) < np.abs(f[i]) < (W_0+f_delta):
            y[i] = (1 + np.cos((np.pi * (abs(f[i]) - W_0 + f_delta)) / (2 * f_delta))) / 2
        else:
            y[i] = 0
    return y

# Sampling rate 1000 hz / second
f = np.linspace(-3, 3, 1000, endpoint=True)

# Plot the square wave signal
plot.plot(f, raised_cosine(f,W_0,f_delta))

# Give a title for the square wave plot
plot.title('Raised Cosine Filter')

# Give x axis label for the square wave plot
plot.xlabel('f')

# Give y axis label for the square wave plot
plot.ylabel('P(f)')
plot.grid(True, which='both')

# Provide x axis and line color
plot.axhline(y=0, color='k')

# Set the max and min values for y axis
plot.ylim(0, 1.2)

# Display the square wave drawn
plot.show()


