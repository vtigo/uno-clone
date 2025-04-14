# Pixel Art Guidelines for UNO Clone

## Overview
Creating pixel art for your UNO clone requires a balance between recognizability and the unique charm of pixel aesthetics. This guide provides direction for creating consistent, visually appealing pixel art assets for your game.

## Canvas Sizes

### Cards
- **Card Size**: 64x96 pixels (4:6 ratio)
- **Card Back**: Same dimensions as cards (64x96)

### UI Elements
- **Buttons**: 96x32 pixels
- **Title Logo**: 320x160 pixels
- **Icons**: 16x16 or 32x32 pixels
- **Color Wheel**: 128x128 pixels

## Color Palette

### Limited Palette Approach
Using a limited color palette will give your game a cohesive look. Aim for approximately 20-32 colors total.

### UNO Colors
Define clear pixel versions of the standard UNO colors:
- **Red**: Primary (#FF0000), Shadow (#C00000), Highlight (#FF6666)
- **Blue**: Primary (#0000FF), Shadow (#0000C0), Highlight (#6666FF)
- **Green**: Primary (#00CC00), Shadow (#009900), Highlight (#66FF66)
- **Yellow**: Primary (#FFCC00), Shadow (#CC9900), Highlight (#FFEE66)
- **Wild/Black**: Primary (#333333), Shadow (#000000), Highlight (#666666)

### UI Colors
- **Background**: Light neutral color (#EEEEEE)
- **Text**: Dark color for contrast (#111111)
- **Accents**: Bright colors that complement but stand out from the UNO colors

## Style Guidelines

### General Style
- Use a consistent level of detail throughout all assets
- Stick to a 45° angle for all shadows
- Limit animation frames to essentials (3-5 frames per animation)
- Keep outlines consistent (either all elements have outlines or none do)

### Card Design
1. **Number Cards**:
   - Large, clear numerals in the center
   - Smaller numeral in top-left and bottom-right corners
   - Card color as the background
   - Simple border design

2. **Action Cards**:
   - Recognizable symbols for Skip, Reverse, and Draw Two
   - Symbol should be larger in the center
   - Smaller symbol in corners like number cards

3. **Wild Cards**:
   - Four-color design using all UNO colors
   - Distinct pattern for Wild vs. Wild Draw Four

4. **Card Back**:
   - Consistent design for all cards
   - Include a simplified game logo or distinct pattern
   - Use contrasting colors that don't match any specific card color

### UI Design
1. **Buttons**:
   - Clear borders and slightly rounded corners
   - Visibly different states (normal, hover, pressed)
   - Consistent padding around text

2. **Title Screen**:
   - Bold, colorful logo incorporating UNO card colors
   - Clean layout with clear hierarchy
   - Animated elements for visual interest (subtle card animations)

3. **Game Screen**:
   - Clean separation between play area and UI elements
   - Clear indications of whose turn it is
   - Distinct area for discard pile vs. draw pile

## Specific Asset Checklist

### Essential Cards (108 total)
- **Number Cards (76)**:
  - 19 Red cards (0-9, with duplicates of 1-9)
  - 19 Blue cards (0-9, with duplicates of 1-9)
  - 19 Green cards (0-9, with duplicates of 1-9)
  - 19 Yellow cards (0-9, with duplicates of 1-9)

- **Action Cards (24)**:
  - 8 Skip cards (2 of each color)
  - 8 Reverse cards (2 of each color)
  - 8 Draw Two cards (2 of each color)

- **Wild Cards (8)**:
  - 4 Wild cards
  - 4 Wild Draw Four cards

### UI Elements
- Title logo
- Play button
- Rules button
- Settings button
- Exit button
- Back button
- UNO button (highlighted)
- End Turn button
- Color wheel with four segments
- Player name display
- Card counter
- Settings toggles

## Pixel Art Techniques

### Shading
- Use **dithering** sparingly for gradients
- Implement **anti-aliasing** for diagonal lines
- Create **outlines** either with darker versions of the fill color or black

### Text
- For small text, use a pixel font (don't try to hand-draw small text)
- For larger text like numbers on cards, custom pixel lettering can work
- Keep a minimum of 5-6 pixels height for legibility

### Special Effects
- Create 3-5 frame animations for:
  - Card dealing
  - Card playing
  - Special card effects
  - UNO button activation
  - Win celebration

## Tools and Workflow

### Recommended Process
1. Start with **rough sketches** of all elements
2. Create **color swatches** to reference throughout
3. Begin with the **most common elements** (basic cards)
4. Build a **template** for each card type, then modify for variations
5. Test assets in-game early to verify sizing and readability

### Organization
- Group files by type (cards, UI, effects)
- Use consistent naming: `cardtype_color_number.png`
- Create separate spritesheets for animations
- Document your color palette for consistency

## Final Tips
- **Consistency** is more important than detail
- Test your designs at the **actual game size**
- Ensure important elements are **readable at a glance**
- Use **animation** sparingly but effectively for important game events
- Remember that **simplicity** often works better for pixel art

Good luck with your pixel art creation! This consistent approach will give your UNO clone a distinctive, cohesive look while maintaining the recognizable elements of the classic card game.
