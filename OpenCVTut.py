import numpy as np
import cv2

def contourDetector():
	# Reading image
	font = cv2.FONT_HERSHEY_COMPLEX
	img2 = cv2.imread('sr71.jpg', cv2.IMREAD_COLOR)
	  
	# Reading same image in another 
	# variable and converting to gray scale.
	img = cv2.imread('sr71.jpg', cv2.IMREAD_GRAYSCALE)
	  
	# Converting image to a binary image
	# ( black and white only image).
	_, threshold = cv2.threshold(img, 110, 255, cv2.THRESH_BINARY)
	  
	# Detecting contours in image.
	contours, _= cv2.findContours(threshold, cv2.RETR_TREE,
	                               cv2.CHAIN_APPROX_SIMPLE)
	  
	# Going through every contours found in the image.
	for cnt in contours :
	  
	    approx = cv2.approxPolyDP(cnt, 0.009 * cv2.arcLength(cnt, True), True)
	  
	    # draws boundary of contours.
	    cv2.drawContours(img2, [approx], 0, (0, 0, 255), 5) 
	  
	    # Used to flatted the array containing
	    # the co-ordinates of the vertices.
	    n = approx.ravel() 
	    i = 0
	  
	    for j in n :
	        if(i % 2 == 0):
	            x = n[i]
	            y = n[i + 1]
	  
	            # String containing the co-ordinates.
	            string = str(x) + " " + str(y) 
	  
	            if(i == 0):
	                # text on topmost co-ordinate.
	                cv2.putText(img2, "Arrow tip", (x, y),
	                                font, 0.5, (255, 0, 0)) 
	            else:
	                # text on remaining co-ordinates.
	                cv2.putText(img2, string, (x, y), 
	                          font, 0.5, (0, 255, 0)) 
	        i = i + 1
		  
	# Showing the final image.
	cv2.imshow('image2', img2) 
	  
	# Exiting the window if 'q' is pressed on the keyboard.
	if cv2.waitKey(0) & 0xFF == ord('q'): 
		cv2.destroyAllWindows()

img = cv2.imread("sr71.jpg", cv2.IMREAD_COLOR) #1, 0 , -1

#split image into color spaces
B, G, R = cv2.split(img)

#add or subtract two images using weights

#Resize
cv2.resize(img, (780, 540), interpolation = cv2.INTER_NEAREST)

#Image erosion
#cv2.erode(src, dst, kernel info)


# Get number of pixel horizontally and vertically.
(height, width) = img.shape[:2]

#cv2.imshow("Cute Kitens", img)

# Specify the size of image along with interploation methods.
# cv2.INTER_AREA is used for shrinking, whereas cv2.INTER_CUBIC
# is used for zooming.
#res = cv2.resize(img, (int(width / 2), int(height / 2)), interpolation = cv2.INTER_CUBIC)

#cv2.imwrite('result.jpg', res)

cv2.imshow('Original Image', img)
cv2.waitKey(0)
  
# Gaussian Blur
Gaussian = cv2.GaussianBlur(img, (7, 7), 0)
#cv2.imshow('Gaussian Blurring', Gaussian)

#cv2.waitKey(0)
  
# Median Blur
median = cv2.medianBlur(img, 5)
#cv2.imshow('Median Blurring', median)
#cv2.waitKey(0)
  
  
# Bilateral Blur
bilateral = cv2.bilateralFilter(img, 9, 75, 75)
#cv2.imshow('Bilateral Blurring', bilateral)
#cv2.waitKey(0) #millisec

imageHSV = cv2.cvtColor(img, cv2.COLOR_BGR2HSV)
cv2.imshow("HSV", imageHSV)
cv2.waitKey(0)
# It is for removing/deleting created GUI window from screen and memory
cv2.destroyAllWindows()

#a 3x3 square black image from scratch
img0 = np.zeros((3, 3), dtype=np.uint8)

import os

# Make an array of 120,000 random bytes.
randomByteArray = bytearray(os.urandom(120000))
flatNumpyArray = np.array(randomByteArray)

# Convert the array to make a 400x300 grayscale image.
grayImage = flatNumpyArray.reshape(300, 400)
cv2.imwrite('RandomGray.png', grayImage)

# Convert the array to make a 400x100 color image.
bgrImage = flatNumpyArray.reshape(100, 400, 3)
cv2.imwrite('RandomColor.png', bgrImage)


# Make an array of 120,000 random bytes.
#randomByteArray = bytearray(os.urandom(120000))
#flatNumpyArray = np.array(randomByteArray)

# Convert the array to make a 400x300 grayscale image.
#grayImage = flatNumpyArray.reshape(300, 400)
#cv2.imwrite('RandomGray.png', grayImage)

# Convert the array to make a 400x100 color image.
#bgrImage = flatNumpyArray.reshape(100, 400, 3)
#cv2.imwrite('RandomColor.png', bgrImage)