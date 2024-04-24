import pygame
import random


def displayControlSticks():

    # Initialize Pygame
    pygame.init()

    # Set up the display
    WIDTH, HEIGHT = 900, 400
    screen = pygame.display.set_mode((WIDTH, HEIGHT))
    pygame.display.set_caption("Dynamic Plot")

    # Define colors
    LIGHT_GREY = (200, 200, 200)
    BLACK = (0, 0, 0)
    RED = (255, 0, 0)
    BLUE = (0, 0, 255)
    ORANGE = (255, 165, 0)

    # Main loop
    clock = pygame.time.Clock()
    running = True
    while running:
        # Handle events
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
        
        # Fill the screen with light grey
        screen.fill(LIGHT_GREY)
        
        # Draw a black border around the screen
        pygame.draw.rect(screen, BLACK, (0, 0, WIDTH, HEIGHT), 1)
        
        # Calculate the center points for the circles
        left_circle_center = (int(WIDTH * 0.25), HEIGHT // 2)
        right_circle_center = (int(WIDTH * 0.75), HEIGHT // 2)
        
        # Draw the circles on the border
        pygame.draw.circle(screen, ORANGE, left_circle_center, 95, 2)
        pygame.draw.circle(screen, ORANGE, right_circle_center, 95, 2)
        
        # Generate random coordinates for the two dots
        dot1_x, dot1_y = int(WIDTH * 0.25), random.randint(0, HEIGHT)
        dot2_x, dot2_y = random.randint(0, WIDTH), HEIGHT //2
        
        # Draw the dots
        pygame.draw.circle(screen, RED, (dot1_x, dot1_y), 5)
        pygame.draw.circle(screen, BLUE, (dot2_x, dot2_y), 5)
        
        # Update the display
        pygame.display.flip()
        
        # Control the refresh rate (increase for higher refresh rates)
        clock.tick(10)

    # Quit Pygame
    pygame.quit()

def eulerToRotaMatrix():
        
    import numpy as np
    import matplotlib.pyplot as plt
    from mpl_toolkits.mplot3d import Axes3D

    # Function to generate rotation matrix from Euler angles
    def euler_to_rotation_matrix(roll, pitch, yaw):
        # Convert Euler angles to radians
        roll = np.radians(roll)
        pitch = np.radians(pitch)
        yaw = np.radians(yaw)

        # Compute rotation matrix
        cy = np.cos(yaw)
        sy = np.sin(yaw)
        cr = np.cos(roll)
        sr = np.sin(roll)
        cp = np.cos(pitch)
        sp = np.sin(pitch)

        R_yaw = np.array([[cy, -sy, 0],
                        [sy, cy, 0],
                        [0, 0, 1]])

        R_pitch = np.array([[cp, 0, sp],
                            [0, 1, 0],
                            [-sp, 0, cp]])

        R_roll = np.array([[1, 0, 0],
                        [0, cr, -sr],
                        [0, sr, cr]])

        R = np.dot(R_yaw, np.dot(R_pitch, R_roll))
        return R

    # Generate trajectory points
    start_position = np.array([0, 0, 0])
    end_position = np.array([10, 5, 5])
    num_points = 5
    trajectory = np.linspace(start_position, end_position, num_points)

    # Generate Euler angles (roll, pitch, yaw) for visualization
    roll = np.linspace(0, -8, num_points)
    pitch = np.linspace(0, 88, num_points)
    yaw = np.linspace(0, 3, num_points)

    # Compute rotation matrices
    rotations = [euler_to_rotation_matrix(roll[i], pitch[i], yaw[i]) for i in range(num_points)]
    #print(rotations)

    # Plot trajectory and drone orientation
    fig = plt.figure()
    ax = fig.add_subplot(111, projection='3d')

    # Plot trajectory
    ax.plot(trajectory[:, 0], trajectory[:, 1], trajectory[:, 2], label='Trajectory')

    # Plot drone orientation
    for i in range(num_points):
        x, y, z = trajectory[i]
        R = rotations[i]
        ax.quiver(x, y, z, R[0, 0], R[1, 0], R[2, 0], color='r', length=1, normalize=True)
        ax.quiver(x, y, z, R[0, 1], R[1, 1], R[2, 1], color='g', length=1, normalize=True)
        ax.quiver(x, y, z, R[0, 2], R[1, 2], R[2, 2], color='b', length=1, normalize=True)

    ax.set_xlabel('X')
    ax.set_ylabel('Y')
    ax.set_zlabel('Z')
    ax.set_title('Drone Trajectory and Orientation')
    plt.legend()
    plt.show()

import pygame
from pygame.locals import *
import math

# Constants
WIDTH, HEIGHT = 800, 600
WHITE = (255, 255, 255)
BLACK = (0, 0, 0)
RED = (255, 0, 0)
GREEN = (0, 255, 0)
BLUE = (0, 0, 255)

# Cube vertices
vertices = [
    (-1, -1, -1),
    (1, -1, -1),
    (1, 1, -1),
    (-1, 1, -1),
    (-1, -1, 1),
    (1, -1, 1),
    (1, 1, 1),
    (-1, 1, 1)
]

# Cube edges
edges = [
    (0, 1), (1, 2), (2, 3), (3, 0),
    (4, 5), (5, 6), (6, 7), (7, 4),
    (0, 4), (1, 5), (2, 6), (3, 7)
]

# Pygame initialization
pygame.init()
screen = pygame.display.set_mode((WIDTH, HEIGHT))
clock = pygame.time.Clock()

# Cube variables
cube_angle_x = 0
cube_angle_y = 0
cube_angle_z = 0

# Main loop
running = True
while running:
    for event in pygame.event.get():
        if event.type == QUIT:
            running = False

    # Handle cube orientation controls
    keys = pygame.key.get_pressed()
    if keys[K_UP]:
        cube_angle_x += 0.01
    if keys[K_DOWN]:
        cube_angle_x -= 0.01
    if keys[K_LEFT]:
        cube_angle_y += 0.01
    if keys[K_RIGHT]:
        cube_angle_y -= 0.01
    if keys[K_q]:
        cube_angle_z += 0.01
    if keys[K_e]:
        cube_angle_z -= 0.01

    # Clear the screen
    screen.fill(BLACK)

    # Project vertices onto 2D surface
    projected_points = []
    for vertex in vertices:
        x = vertex[0]
        y = vertex[1]
        z = vertex[2]

        # Rotate around X axis
        y = y * math.cos(cube_angle_x) - z * math.sin(cube_angle_x)
        z = y * math.sin(cube_angle_x) + z * math.cos(cube_angle_x)

        # Rotate around Y axis
        x = x * math.cos(cube_angle_y) + z * math.sin(cube_angle_y)
        z = -x * math.sin(cube_angle_y) + z * math.cos(cube_angle_y)

        # Rotate around Z axis
        x = x * math.cos(cube_angle_z) - y * math.sin(cube_angle_z)
        y = x * math.sin(cube_angle_z) + y * math.cos(cube_angle_z)

        # Perspective projection
        f = 200 / (z + 5)
        projected_x = x * f + WIDTH / 2
        projected_y = y * f + HEIGHT / 2
        projected_points.append((projected_x, projected_y))

    # Draw edges
    for edge in edges:
        pygame.draw.line(screen, WHITE, projected_points[edge[0]], projected_points[edge[1]], 1)

    pygame.display.flip()
    clock.tick(30)

pygame.quit()
