import matplotlib.pyplot as plot
import numpy as np

T_b=0.2

# Define the p(t) fuction
def p(t):
    return np.sinc(t/(T_b))
    
# Sampling rate 1000 hz / second
t = np.linspace(-1, 1, 1000, endpoint=True)
    
# Plot the square wave signal
plot.plot(t, p(t)+(0.8*p(t-T_b))+(1.2*p(t-2*T_b))+(0.3*p(t-3*T_b)))

# Plot text when t = 0,T_b, 2T_b, 3T_b
tmp = [0, 0.2, 0.4, 0.6]
for i in 0,1,2,3:
    plot.plot(tmp[i], p(tmp[i])+(0.8*p(tmp[i]-T_b))+(1.2*p(tmp[i]-2*T_b))+(0.3*p(tmp[i]-3*T_b)), 'ro') # as red dots
    val = (p(tmp[i])+(0.8*p(tmp[i]-T_b))+(1.2*p(tmp[i]-2*T_b))+(0.3*p(tmp[i]-3*T_b)))
    plot.text(tmp[i], val, val.astype(str), fontsize=8)

# Give a title for the square wave plot
plot.title('combined of sinc function')

# Give x axis label for the square wave plot
plot.xlabel('t')

# Give y axis label for the square wave plot
plot.ylabel('p(t)')
plot.grid(True, which='both')

# Provide x axis and line color
plot.axhline(y=0, color='k')

# Set the max and min values for y axis
plot.ylim(-1.3, 1.3)

# Display the square wave drawn
plot.show()