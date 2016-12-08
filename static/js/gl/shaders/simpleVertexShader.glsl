attribute vec3 vertexPosition;
attribute vec2 textureCoordinate;

uniform mat4 modelViewMatrix;
uniform mat4 perspectiveMatrix;
uniform mat4 cameraRotationMatrix;

varying vec2 textureCoordinateVarying;

void main(void) {
  gl_Position = perspectiveMatrix * cameraRotationMatrix * modelViewMatrix * vec4(vertexPosition, 1.0);
  textureCoordinateVarying = textureCoordinate;
}
