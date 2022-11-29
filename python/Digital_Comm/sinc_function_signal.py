import matplotlib.pyplot as plot
import numpy as np

T_b=0.2
# Sampling rate 1000 hz / second
t = np.linspace(0, 1.5, 1000, endpoint=True)

# Define the sinc function
def sinc(x):
    return np.sin(x)/x

# Define the p(t) fuction
def p(t):
    return sinc(t/T_b)

# Plot the square wave signal
plot.plot(t, p(t)+(0.8*p(t-T_b))+(1.2*p(t-2*T_b))+(0.3*p(t-3*T_b)))

# Give a title for the square wave plot
plot.title('combined of sinc function')

# Give x axis label for the square wave plot
plot.xlabel('t')

# Give y axis label for the square wave plot
plot.ylabel('y')
plot.grid(True, which='both')

# Provide x axis and line color
plot.axhline(y=0, color='k')

# Set the max and min values for y axis
plot.ylim(-6, 6)

# Display the square wave drawn
plot.show()