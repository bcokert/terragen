"use strict";

const RenderComponent = require("../render-component");

const ColorableProto = {
    /**
     * Changes the color of the object
     * @param {Float32Array} newColor - A vec4 describing the new color
     * @returns {undefined}
     */
    setColor(newColor) {
        this.color = newColor;
    }
};

/**
 * Adds the ability for an object to be colored.
 * @param {Float32Array} color
 * @constructor
 * @returns {Object}
 */
const Colorable = (color) => {
    if (Object.prototype.toString.call(color) !== "[object Float32Array]" || color.length != 4) {
        throw new TypeError("Invalid color provided to RenderObject factory.");
    }

    return RenderComponent({
        name: "Colorable",
        proto: ColorableProto,
        properties: {
            color: color
        },
        renderFn: (instance, gl, programInfo, modelViewMatrix) => {
            programInfo.uniforms.inputColor.set(instance.color);
            return modelViewMatrix;
        }
    });
};

Colorable.prototype = ColorableProto;

module.exports = Colorable;
