"use strict";

const RenderComponent = require("../render-component");

const TexturableProto = {
    /**
     * Binds and buffers uv data
     * @param {WebGLRenderingContext} gl
     * @returns {undefined}
     */
    updateUVArray(gl) {
        gl.bindBuffer(gl.ARRAY_BUFFER, this.uvBuffer);
        gl.bufferData(gl.ARRAY_BUFFER, this.uvArray, gl.STATIC_DRAW);
    }
};

/**
 * Adds the ability for an object to have a renderable triangle mesh.
 * @param {WebGLRenderingContext} gl
 * @param {Float32Array} uvArray - an array containing uv coordinates for this object
 * @param {WebGLTexture|null} texture
 * @constructor
 * @returns {Object}
 */
const Texturable = (gl, uvArray, texture) => {
    if (Object.prototype.toString.call(uvArray) !== "[object Float32Array]") {
        throw new TypeError("Invalid uvArray provided to Texturable factory.");
    }

    return RenderComponent({
        name: "Texturable",
        proto: TexturableProto,
        properties: {
            uvArray: uvArray,
            uvBuffer: gl.createBuffer(),
            texture: texture
        },
        initializer: instance => {
            instance.updateUVArray(gl);
        },
        renderFn: (instance, gl, programInfo, modelViewMatrix) => {
            if (gl.isTexture(instance.texture)) {
                programInfo.uniforms.texture1.set(instance.texture);
                programInfo.attributes.textureCoordinate.set(instance.uvBuffer);
            }
            return modelViewMatrix;
        }
    });
};

Texturable.prototype = TexturableProto;

module.exports = Texturable;
