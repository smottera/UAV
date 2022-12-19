from exif import Image

img_path2 = f'C:/Users/Juggernaut/Desktop/PANcard.jpg'
img_path = f'C:/Users/Juggernaut/Desktop/DJIMini3Pro/DJI_0135.JPG'


with open(img_path, 'rb') as img_file:
    img = Image(img_file)
    
print(img.has_exif)

count = 0
for item in sorted(img.list_all()):
    print(item)
    count += 1

print("Total number of keys: ", count)

print(
    f'Model: {img.get("model")}'
    f'colorSpace: {img.get("color_space")}',
    f'contrast: {img.get("contrast")}'
    f'aperture_value: {img.get("aperture_value")}',
    f'DzoomRatio: {img.get("digital_zoom_ratio")}'
    f'fNumber: {img.get("f_number")}',
    

)