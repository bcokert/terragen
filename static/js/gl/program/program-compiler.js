"use strict";

const ProgramCompilerProto = {
    /**
     * Takes the source of a glsl shader and compiles it, returning the resulting WebGLShader, or an Error if it fails
     * @param {WebGLRenderingContext} gl The gl rendering context to compile with
     * @param {string} shaderSource The source of the shader, typically acquired via require('../path/to/shader.xx')
     * @param {number} shaderType The shader type, either gl.FRAGMENT_SHADER or gl.VERTEX_SHADER.
     * @throws {TypeError}
     * @returns {WebGLShader|Error}
     */
    compileShader(gl, shaderSource, shaderType) {
        if (typeof shaderSource !== "string" || shaderSource === "") {
            throw new TypeError("Invalid shaderSource provided to compileShader.");
        }
        if (shaderType !== gl.VERTEX_SHADER && shaderType !== gl.FRAGMENT_SHADER) {
            throw new TypeError("Invalid shaderType provided to compileShader.");
        }
        if (!gl || !gl.createShader) {
            throw new TypeError("Invalid rendering context provided to compileShader.");
        }

        var shader = gl.createShader(shaderType);
        gl.shaderSource(shader, shaderSource);
        gl.compileShader(shader);

        if (!gl.getShaderParameter(shader, gl.COMPILE_STATUS)) {
            return new Error("Unable to compile shader: " + gl.getShaderInfoLog(shader));
        }

        return shader;
    },

    /**
     * Combines compiled shaders into a WebGLProgram and returns it, or an Error if it fails
     * @param {WebGLRenderingContext} gl The gl rendering context to compile with
     * @param {WebGLShader} vertexShader The compiled vertex shader, eg that from compileShader()
     * @param {WebGLShader} fragmentShader The compiled fragment shader, eg that from compileShader()
     * @returns {WebGLProgram|Error}
     */
    createShaderProgram(gl, vertexShader, fragmentShader) {
        if (!gl || !gl.createProgram) {
            throw new TypeError("Invalid rendering context provided to createShaderProgram.");
        }
        if (!gl.isShader(vertexShader)) {
            throw new TypeError("Invalid vertexShader provided to createShaderProgram.");
        }
        if (!gl.isShader(fragmentShader)) {
            throw new TypeError("Invalid fragmentShader provided to createShaderProgram.");
        }

        var shaderProgram = gl.createProgram();
        gl.attachShader(shaderProgram, vertexShader);
        gl.attachShader(shaderProgram, fragmentShader);
        gl.linkProgram(shaderProgram);

        if (!gl.getProgramParameter(shaderProgram, gl.LINK_STATUS)) {
            return new Error("Unable to create shader program: " + gl.getProgramInfoLog(shaderProgram));
        }

        return shaderProgram;
    }
};

/**
 * Creates a new program compiler, which can create programs from shaders, and compile shaders from sources
 * @returns {ProgramCompiler}
 * @constructor
 */
const ProgramCompiler = () => {
    return Object.assign(Object.create(ProgramCompilerProto), {
        banana: 1
    });
};

ProgramCompiler.prototype = ProgramCompilerProto;

module.exports = ProgramCompiler;
