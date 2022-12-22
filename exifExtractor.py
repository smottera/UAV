from exif import Image

img_path2 = f'C:/Users/Juggernaut/Desktop/PANcard.jpg'
img_path = f'C:/Users/QuNuKhan/Desktop/Mini3Pro/DJI_0155.JPG'


with open(img_path, 'rb') as img_file:
    img = Image(img_file)
    
print(img.has_exif)

count = 0
for item in sorted(img.list_all()):
    #print(item,": ", img.get(item))
    count += 1

print("Total number of keys: ", count)

#pip install python-xmp-toolkit

f = r'C:/Users/QuNuKhan/Desktop/Mini3Pro/DJI_0155.JPG'
fd = open(f, encoding = 'latin-1')
d= fd.read()
xmp_start = d.find('<x:xmpmeta')
xmp_end = d.find('</x:xmpmeta')
xmp_str = d[xmp_start:xmp_end+12]
print(xmp_str)
