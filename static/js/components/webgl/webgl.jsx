"use strict";
var React = require("react");
var SimpleVertexShaderSource = require("../../gl/shaders/simpleVertexShader.glsl");
var SimpleFragmentShaderSource = require("../../gl/shaders/simpleFragmentShader.glsl");
var SimpleShaderInfo = require("../../gl/shaders/simpleShaderInfo");
var GLMatrix = require("gl-matrix");
var ResourceLoader = require("../../gl/resource/resource-loader");
var InputManager = require("../../gl/input/input-manager");
var ProgramManager = require("../../gl/program/program-manager");
var ProgramCompiler = require("../../gl/program/program-compiler");
var Cube = require("../../gl/render-objects/cube");

require("./webgl.less");

class WebGL extends React.Component {
    constructor(props) {
        super(props);

        this.canvas = null; // ref to root canvas element

        this.state = {
            frameRate: 0,
            initError: undefined,
            isLoadingTextures: false
        };
    }

    componentDidMount() {
        var gl = this.createGLContext(this.canvas);
        this.setState({isLoadingTextures: true});
        if (gl) {
            const resourceLoader = ResourceLoader();
            resourceLoader.loadResources(gl).then(resources => {
                this.setState({isLoadingTextures: false});
                // Configure some gl globals
                gl.clearColor(0.0, 0.0, 0.0, 1.0);
                gl.enable(gl.DEPTH_TEST);

                // Create the program manager
                var programManager = ProgramManager(gl, {
                    simple: {
                        vertexShader: SimpleVertexShaderSource,
                        fragmentShader: SimpleFragmentShaderSource,
                        programInfo: SimpleShaderInfo
                    }
                }, ProgramCompiler());

                // Create the perspective matrix
                var perspectiveMatrix = GLMatrix.mat4.create();
                GLMatrix.mat4.perspective(perspectiveMatrix, 45, gl.drawingBufferWidth / gl.drawingBufferHeight, 0.1, 100.0);

                // Create the world modelView matrix
                var modelViewMatrix = GLMatrix.mat4.create();

                // Create the camera rotation matrix
                var cameraRotationMatrix = GLMatrix.mat4.create();

                // Create global input manager
                var inputManager = InputManager();

                // Add some event handlers for the mouse, to modify the rotation matrix
                var isMouseDragging = false;
                var mouseDragStartCoords = null;
                inputManager.on("mousedown", (inputState, coords) => {
                    isMouseDragging = true;
                    mouseDragStartCoords = coords;
                });

                inputManager.on("mouseup", () => {
                    isMouseDragging = false;
                    mouseDragStartCoords = null;
                });

                inputManager.on("mousemove", (inputState, coords) => {
                    if (isMouseDragging && mouseDragStartCoords) {
                        if (!inputState[InputManager.KEY_TO_CODE["shift"]]) {
                            var dx = mouseDragStartCoords.x - coords.x;
                            var dy = mouseDragStartCoords.y - coords.y;
                            GLMatrix.mat4.rotate(cameraRotationMatrix, cameraRotationMatrix, dx / 300, [0, 1, 0]);
                            GLMatrix.mat4.rotate(cameraRotationMatrix, cameraRotationMatrix, dy / 300, [1, 0, 0]);
                            mouseDragStartCoords.x = coords.x;
                            mouseDragStartCoords.y = coords.y;
                        }
                    }
                });

                // Create test cubes
                var cubes = [1, 1, 1].map((size, i) => Cube(
                    gl,
                    new Float32Array([(i - 1) * 5, i - 1, -8]),
                    GLMatrix.vec4.fromValues(Math.random(), Math.random(), Math.random(), 1),
                    resources.textures["testWoodTexture"],
                    inputManager
                ));

                // Start the system, which starts the render loop and logic loop
                this.start(gl, programManager, perspectiveMatrix, modelViewMatrix, cameraRotationMatrix, cubes);
            });
        } else {
            this.setState({initError: "Unable to create context. The browser does not support WebGL."});
        }
    }

    /**
     * Creates the main gl context object
     * @param {HTMLElement} canvas
     * @returns {WebGLRenderingContext}
     */
    createGLContext(canvas) {
        var gl;
        try {
            gl = canvas.getContext("webgl") || canvas.getContext("experimental-webgl");
        } catch (e) {}

        return gl ? gl : undefined;
    }

    /**
     * Takes an initialized rendering context and all the uniforms, attributes, and objects it needs, then
     * kicks off the renderLoop and logicLoop
     * @param {WebGLRenderingContext} gl
     * @param {ProgramManager} programManager
     * @param {Float32Array} perspectiveMatrix
     * @param {Float32Array} modelViewMatrix
     * @param {Float32Array} cameraRotationMatrix
     * @param {Cube[]} cubes
     * @returns {undefined}
     */
    start(gl, programManager, perspectiveMatrix, modelViewMatrix, cameraRotationMatrix, cubes) {
        var renderLoop = this.renderLoopFactory(gl, programManager, perspectiveMatrix, modelViewMatrix, cameraRotationMatrix, cubes);
        var logicLoop = this.logicLoopFactory(cubes);

        renderLoop();
        logicLoop();
    }

    /**
     * Creates a renderLoop from the rendering context and everything required to render
     * The resulting render loop will run as fast as the underlying hardware will allow
     * @param {WebGLRenderingContext} gl
     * @param {ProgramManager} programManager
     * @param {Float32Array} perspectiveMatrix
     * @param {Float32Array} modelViewMatrix
     * @param {Float32Array} cameraRotationMatrix
     * @param {Cube[]} cubes
     * @returns {function} The function to call to start the render loop
     */
    renderLoopFactory(gl, programManager, perspectiveMatrix, modelViewMatrix, cameraRotationMatrix, cubes) {
        var lastRenderTime = 0;
        var msPerFrame = 0;
        var now = 0;
        var renderLoop = () => {
            requestAnimationFrame(renderLoop);

            // Update the frame rate
            now = Date.now();
            msPerFrame = now - lastRenderTime;
            lastRenderTime = now;

            // Update the viewport config if it has changed
            if (gl.canvas.width !== gl.canvas.clientWidth || gl.canvas.height !== gl.canvas.clientHeight) {
                gl.canvas.width = gl.canvas.clientWidth;
                gl.canvas.height = gl.canvas.clientHeight;
            }

            // Reset the viewport to the new config, and clear the render buffer
            gl.viewport(0, 0, gl.drawingBufferWidth, gl.drawingBufferHeight);
            gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);

            // Get the main program
            const mainProgram = programManager.getProgram("simple");
            gl.useProgram(mainProgram.program);
            const programInfo = mainProgram.programInfo;

            // Push the perspective matrix
            programInfo.uniforms.perspectiveMatrix.set(perspectiveMatrix);

            // Push the camera rotation matrix
            programInfo.uniforms.cameraRotationMatrix.set(cameraRotationMatrix);

            // Render the objects
            cubes.forEach(cube => cube.render(gl, programInfo, GLMatrix.mat4.clone(modelViewMatrix)));
        };

        return () => {
            lastRenderTime = Date.now();
            setInterval(() => this.setState({frameRate: parseInt(1 / (msPerFrame / 1000), 10)}), 5000);
            renderLoop();
        };
    }

    /**
     * Creates a logicLoop for the current scene
     * The resulting logicLoop runs at a constant rate, unless the cpu is actually overloaded (which is rare since webGL will prefer to drop render calls instead)
     * It will pass in the delta time to all children to handle this case
     * @param {Cube[]} cubes
     * @returns {function} The function call to start the logic loop
     */
    logicLoopFactory(cubes) {
        var now = 0;
        var lastTime = 0;
        var delta = 0;

        var logicFunction = () => {
            now = Date.now();
            delta = (now - lastTime)/1000;
            lastTime = now;

            cubes.forEach(cube => cube.update(delta));
        };

        return () => {
            lastTime = Date.now();
            setInterval(logicFunction, 1000/60);
        };
    }

    render() {
        if (this.state.initError) {
            return <p>{this.state.initError}</p>;
        }

        var notice = null;
        if (this.state.initError) {
            notice = <p>{this.state.initError}</p>;
        }
        if (this.state.isLoadingTextures) {
            notice = <p>Please wait while we load textures</p>;
        }
        return (
            <div className="WebGL">
                {notice}
                <div className="-frameRateDisplay">Frame Rate (±3): {this.state.frameRate}</div>
                <canvas key="renderCanvas" ref={node => this.canvas = node}/>
            </div>
        );
    }
}

WebGL.displayName = "WebGL";

WebGL.propTypes = {};

module.exports = WebGL;
