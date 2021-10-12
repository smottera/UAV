import numpy as np
import cv2

img = cv2.imread("sr71.jpg", cv2.IMREAD_COLOR) #1, 0 , -1

#split image into color spaces
B, G, R = cv2.split(img)

#add or subtract two images using weights

#Resize
cv2.resize(img, (780, 540), interpolation = cv2.INTER_NEAREST)

# Get number of pixel horizontally and vertically.
(height, width) = img.shape[:2]

cv2.imshow("Cute Kitens", img)

# Specify the size of image along with interploation methods.
# cv2.INTER_AREA is used for shrinking, whereas cv2.INTER_CUBIC
# is used for zooming.
res = cv2.resize(img, (int(width / 2), int(height / 2)), interpolation = cv2.INTER_CUBIC)

#cv2.imwrite('result.jpg', res)

cv2.waitKey(0) #millisec
 
# It is for removing/deleting created GUI window from screen and memory
cv2.destroyAllWindows()
