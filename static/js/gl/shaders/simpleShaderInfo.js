"use strict";

/**
 * @param {WebGLRenderingContext} gl
 * @param {WebGLProgram} program
 * @returns {ProgramInfo}
 */
module.exports = function(gl, program) {
    // Vertex Attributes
    const vertexPosition = gl.getAttribLocation(program, "vertexPosition");
    gl.enableVertexAttribArray(vertexPosition);
    const textureCoordinate = gl.getAttribLocation(program, "textureCoordinate");
    let textureCoordinateEnabled = false;

    // Vertex Uniforms
    const cameraRotationMatrix = gl.getUniformLocation(program, "cameraRotationMatrix");
    const perspectiveMatrix = gl.getUniformLocation(program, "perspectiveMatrix");
    const modelViewMatrix = gl.getUniformLocation(program, "modelViewMatrix");

    // Fragment Uniforms
    const inputColor = gl.getUniformLocation(program, "inputColor");
    const texture1 = gl.getUniformLocation(program, "texture1");

    return {
        attributes: {
            vertexPosition: {
                location: vertexPosition,
                set: vertexBuffer => {
                    gl.bindBuffer(gl.ARRAY_BUFFER, vertexBuffer);
                    gl.vertexAttribPointer(vertexPosition, 3, gl.FLOAT, false, 0, 0);
                }
            },
            textureCoordinate: {
                location: textureCoordinate,
                set: textureCoordinateBuffer => {
                    if (!textureCoordinateEnabled) {
                        gl.enableVertexAttribArray(textureCoordinate);
                        textureCoordinateEnabled = true;
                    }
                    gl.bindBuffer(gl.ARRAY_BUFFER, textureCoordinateBuffer);
                    gl.vertexAttribPointer(textureCoordinate, 2, gl.FLOAT, false, 0, 0);
                },
                disable: () => {
                    gl.disableVertexAttribArray(textureCoordinate);
                    textureCoordinateEnabled = false;
                }
            }
        },
        uniforms: {
            cameraRotationMatrix: {
                location: cameraRotationMatrix,
                /**
                 * @param {Float32Array} matrix - 4x4
                 * @returns {undefined}
                 */
                set: matrix => {
                    gl.uniformMatrix4fv(cameraRotationMatrix, false, matrix);
                }
            },
            perspectiveMatrix: {
                location: perspectiveMatrix,
                /**
                 * @param {Float32Array} matrix - 4x4
                 * @returns {undefined}
                 */
                set: matrix => {
                    gl.uniformMatrix4fv(perspectiveMatrix, false, matrix);
                }
            },
            modelViewMatrix: {
                location: modelViewMatrix,
                /**
                 * @param {Float32Array} matrix - 4x4
                 * @returns {undefined}
                 */
                set: matrix => {
                    gl.uniformMatrix4fv(modelViewMatrix, false, matrix);
                }
            },
            inputColor: {
                location: inputColor,
                /**
                 * @param {Float32Array} color - 4x1
                 * @returns {undefined}
                 */
                set: color => {
                    gl.uniform4fv(inputColor, color);
                }
            },
            texture1: {
                location: texture1,
                /**
                 * @param {WebGLTexture} texture
                 * @returns {undefined}
                 */
                set: texture => {
                    gl.uniform1i(texture1, 0);
                    gl.activeTexture(gl.TEXTURE0);
                    gl.bindTexture(gl.TEXTURE_2D, texture);
                }
            }
        }
    };
};
