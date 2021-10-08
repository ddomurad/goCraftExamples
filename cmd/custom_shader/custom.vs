#version 330 core

layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTex;

uniform mat4 uTrans;
uniform mat4 uProj;
uniform mat4 uView;

uniform float animation;

out vec2 texCoord;

void main(){
    
    float angle = atan(aPos.y, aPos.x);
    float xDir = cos(angle);
    float yDir = sin(angle);
    float boundedAnimation = sin(animation/5)*20;

    vec3 offset = vec3(xDir * sin(angle*boundedAnimation), yDir* sin(angle*boundedAnimation),0)/10;
    //aPos.y = sin(angle);
    vec4 pos = uProj * uView * uTrans * vec4(aPos + offset, 1.0);
    
    gl_Position = pos;
    texCoord = aTex;
}