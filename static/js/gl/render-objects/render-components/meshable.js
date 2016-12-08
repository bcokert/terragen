"use strict";

const RenderComponent = require("../render-component");

const MeshableProto = {
    /**
     * Binds and buffers vertex data
     * @param {WebGLRenderingContext} gl
     * @returns {undefined}
     */
    updateVertexBuffer(gl) {
        gl.bindBuffer(gl.ARRAY_BUFFER, this.vertexBuffer);
        gl.bufferData(gl.ARRAY_BUFFER, this.vertexArray, gl.STATIC_DRAW);
    },

    /**
     * Binds and buffers index data
     * @param {WebGLRenderingContext} gl
     * @returns {undefined}
     */
    updateIndexBuffer(gl) {
        gl.bindBuffer(gl.ELEMENT_ARRAY_BUFFER, this.indexBuffer);
        gl.bufferData(gl.ELEMENT_ARRAY_BUFFER, this.indexArray, gl.STATIC_DRAW);
    }
};

/**
 * Adds the ability for an object to have a renderable triangle mesh.
 * @param {WebGLRenderingContext} gl
 * @param {Float32Array} vertexArray - an array containing triples of floats, each representing a vertex
 * @param {Uint16Array} indexArray - an array containing triples of indexes, each pointing to 3 vertices that make up a triangle
 * @constructor
 * @returns {Object}
 */
const Meshable = (gl, vertexArray, indexArray) => {
    if (Object.prototype.toString.call(vertexArray) !== "[object Float32Array]") {
        throw new TypeError("Invalid vertexArray provided to Meshable factory.");
    }
    if (Object.prototype.toString.call(indexArray) !== "[object Uint16Array]") {
        throw new TypeError("Invalid indexArray provided to Meshable factory.");
    }

    return RenderComponent({
        name: "Meshable",
        initializer: instance => {
            instance.updateVertexBuffer(gl);
            instance.updateIndexBuffer(gl);
        },
        proto: MeshableProto,
        properties: {
            vertexArray: vertexArray,
            vertexBuffer: gl.createBuffer(),
            indexArray: indexArray,
            indexBuffer: gl.createBuffer()
        },
        renderFn: (instance, gl, programInfo, modelViewMatrix) => {
            programInfo.attributes.vertexPosition.set(instance.vertexBuffer);
            return modelViewMatrix;
        },
        finishRenderFn: (instance, gl, programInfo, modelViewMatrix) => {
            programInfo.uniforms.modelViewMatrix.set(modelViewMatrix);
            gl.drawElements(gl.TRIANGLES, instance.indexArray.length, gl.UNSIGNED_SHORT, 0);
        }
    });
};

module.exports = Meshable;
