"use strict";

const RenderObject = require("./render-object");
const RenderComponent = require("./render-component");

const Rotatable = require("./render-components/rotateable");
const Positionable = require("./render-components/positionable");
const Meshable = require("./render-components/meshable");
const Texturable = require("./render-components/texturable");
const Colorable = require("./render-components/colorable");
const Movable = require("./render-components/movable");

const InputManager = require("../input/input-manager");
const CubeMesh = require("../prebuilt-meshes/cube");

/**
 * A cube that can be rendered, rotated, textured, colored, and moved.
 * Contains some default controls
 * @param {WebGLRenderingContext} gl
 * @param {Float32Array} position
 * @param {Float32Array} color
 * @param {WebGLTexture|null} texture
 * @param {InputManager} inputManager
 * @returns {Object}
 * @constructor
 */
const Cube = (gl, position, color, texture, inputManager) => {

    let instance = RenderObject([
        Rotatable(new Float32Array([0, 0, 0])),
        Positionable(position),
        Meshable(gl, CubeMesh.vertexArray, CubeMesh.indexArray),
        Texturable(gl, CubeMesh.uvArray, texture),
        Colorable(color),
        Movable(8)
    ]);

    const handleInput = inputState => {
        instance.acceleration[0] = (inputState[InputManager.KEY_TO_CODE["A"]] ? -10 : 0) + (inputState[InputManager.KEY_TO_CODE["D"]] ? 10 : 0);
        instance.acceleration[1] = (inputState[InputManager.KEY_TO_CODE["S"]] ? -10 : 0) + (inputState[InputManager.KEY_TO_CODE["W"]] ? 10 : 0);

        if (inputState[InputManager.KEY_TO_CODE[" "]]) {
            instance.acceleration[0] = 0;
            instance.acceleration[1] = 0;
            instance.velocity[0] = 0;
            instance.velocity[1] = 0;
        }

        let rx = (inputState[InputManager.KEY_TO_CODE["up"]] ? -.1 : .1) + (inputState[InputManager.KEY_TO_CODE["down"]] ? .1 : -.1);
        let ry = (inputState[InputManager.KEY_TO_CODE["left"]] ? .1 : -.1) + (inputState[InputManager.KEY_TO_CODE["right"]] ? -.1 : .1);
        instance.rotate([rx, ry, 0]);
    };

    inputManager.on("keydown", handleInput);
    inputManager.on("keyup", handleInput);

    return instance;
};

module.exports = Cube;
