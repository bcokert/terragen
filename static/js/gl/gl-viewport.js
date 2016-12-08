"use strict";

const GLViewportProto = {
    start: () => {

    }
};

/**
 * Creates a new GLViewport, which is passed the canvas that will be used.
 * The viewport handles rendering and updating all of its content, but it does not naturally interact with the DOM
 * @param {HTMLCanvasElement} canvas
 * @returns {GLViewPort}
 * @constructor
 */
const GLViewport = (canvas) => {
    let gl;
    try {
        gl = canvas.getContext("webgl") || canvas.getContext("experimental-webgl");
    } catch (e) {
        return Error("This browser does not support webgl");
    }

    return Object.assign(Object.create(GLViewportProto), {
        /**
         * @type {WebGLRenderingContext} gl
         */
        gl: gl
    });
};

GLViewport.prototype = GLViewportProto;

module.exports = GLViewport;
