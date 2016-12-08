precision mediump float;

uniform vec4 inputColor;
uniform sampler2D texture1;

varying vec2 textureCoordinateVarying;

void main(void) {
    vec4 textureColor = texture2D(texture1, vec2(textureCoordinateVarying.s, textureCoordinateVarying.t));
    gl_FragColor = textureColor * inputColor;
}
