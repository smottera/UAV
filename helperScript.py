import handRecog
import numpy as np
# Color histogram
from matplotlib import pyplot as plt
import cv2 as cv

img = cv.imread('quarry.JPG')

def histogrammer(img):
    chans = cv.split(img)
    colors = 'b', 'g', 'r'

    plt.figure()
    plt.title('Flattened color histogram')
    plt.xlabel('Bins')
    plt.ylabel('# of pixels')

    for (chan, color) in zip(chans, colors):
        hist = cv.calcHist([chan], [0], None, [256], [0, 255])
        plt.plot(hist, color=color)
        plt.xlim([0, 256])
        plt.ylim([0, 172000])

    plt.show()
    cv.waitKey(0)



histogrammer(img)