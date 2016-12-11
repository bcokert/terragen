"use strict";
var React = require("react");
var ReactDom = require("react-dom");
var NoiseBrowserList = require("./components/noise-browser-list/noise-browser-list.jsx");
var Button = require("./components/control/button/button.jsx");
var WebGL = require("./components/webgl/webgl");

require("./index.less");

class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            page: "Main"
        };

        this.changePage = this.changePage.bind(this);
        this.onPageReload = this.onPageReload.bind(this);
        this.onPageRefresh = this.onPageRefresh.bind(this);
    }

    componentDidMount() {
        window.addEventListener("popstate", this.onPageReload);
        window.onload = this.onPageRefresh;
    }

    onPageRefresh() {
        this.setState({page: history.state && history.state.page ? history.state.page : "Main"});
    }

    onPageReload(e) {
        if (!e || typeof e !== "object" || !e.state) {
            this.setState({page: "Main"});
        } else {
            this.setState({page: e.state.page});
        }
    }

    changePage(newPage) {
        history.pushState({page: newPage}, newPage, "/");
        this.setState({
            page: newPage
        });
    }

    renderMainPage() {
        return (
            <div className="-mainPage">
                <div className="-column">
                    <p className="-title">Spectral 1D</p>
                    <Button onClick={() => this.changePage("Spectral1D")}>Spectral Noise 1D</Button>
                </div>
                <div className="-column">
                    <p className="-title">Spectral 2D</p>
                    <Button onClick={() => this.changePage("Spectral2D")}>Spectral Noise 2D</Button>
                </div>
                <div className="-column">
                    <p className="-title">Lattice 2D</p>
                    <Button onClick={() => this.changePage("Lattice2D")}>Lattice Noise 2D</Button>
                </div>
                <div className="-column">
                    <p className="-title">WebGl</p>
                    <Button onClick={() => this.changePage("WebGLDemo")}>WebGL Demo</Button>
                </div>
            </div>
        );
    }

    renderSpectral1DList() {
        return (
            <NoiseBrowserList
                browserList={[{
                    dimension: 1,
                    displayName: "Red Noise",
                    endpoint: "/noise",
                    noiseFunction: "red"
                },{
                    dimension: 1,
                    displayName: "Pink Noise",
                    endpoint: "/noise",
                    noiseFunction: "pink"
                },{
                    dimension: 1,
                    displayName: "White Noise",
                    endpoint: "/noise",
                    noiseFunction: "white"
                },{
                    dimension: 1,
                    displayName: "Blue Noise",
                    endpoint: "/noise",
                    noiseFunction: "blue"
                },{
                    dimension: 1,
                    displayName: "Violet Noise",
                    endpoint: "/noise",
                    noiseFunction: "violet"
                }]}
                description="Spectral Noise is created from random sinusoids of various frequencies, combined via a weighted sum, where the weights are related to the frequency"
                generator="Random"
                presetCollectionName="Spectral Noise"
                synthesizer="Octave"
                transformer="Sinusoid"
            />
        );
    }

    renderSpectral2DList() {
        return (
            <NoiseBrowserList
                browserList={[{
                    dimension: 2,
                    displayName: "Red Noise",
                    endpoint: "/noise",
                    noiseFunction: "red"
                },{
                    dimension: 2,
                    displayName: "Pink Noise",
                    endpoint: "/noise",
                    noiseFunction: "pink"
                },{
                    dimension: 2,
                    displayName: "White Noise",
                    endpoint: "/noise",
                    noiseFunction: "white"
                },{
                    dimension: 2,
                    displayName: "Blue Noise",
                    endpoint: "/noise",
                    noiseFunction: "blue"
                },{
                    dimension: 2,
                    displayName: "Violet Noise",
                    endpoint: "/noise",
                    noiseFunction: "violet"
                },{
                    dimension: 2,
                    displayName: "Raw Perlin Noise",
                    endpoint: "/noise",
                    noiseFunction: "rawPerlin"
                }]}
                description="Spectral Noise is created from random sinusoids of various frequencies, combined via a weighted sum, where the weights are related to the frequency"
                generator="Random"
                presetCollectionName="Spectral Noise"
                synthesizer="Octave"
                transformer="Sinusoid"
            />
        );
    }

    renderLattice2DList() {
        return (
            <NoiseBrowserList
                browserList={[{
                    dimension: 2,
                    displayName: "Raw Perlin Noise",
                    endpoint: "/noise",
                    noiseFunction: "rawPerlin"
                }]}
                description="Lattice noise is created by starting with a base n dimensional grid (or lattice) of noisy data, then interpolating values between grid points."
                generator="Perlin"
                presetCollectionName="Lattice Noise"
                synthesizer="N/A"
                transformer="N/A"
            />
        );
    }

    renderWebGLDemo() {
        return <WebGL />;
    }

    render() {
        let page;
        switch (this.state.page) {
            case "Main":
                page = this.renderMainPage();
                break;
            case "Spectral1D":
                page = this.renderSpectral1DList();
                break;
            case "Spectral2D":
                page = this.renderSpectral2DList();
                break;
            case "Lattice2D":
                page = this.renderLattice2DList();
                break;
            case "WebGLDemo":
                page = this.renderWebGLDemo();
                break;
            default:
                page = null;
        }

        return (
            <div className="App">
                {page}
            </div>
        );
    }
}

App.displayName = "App";

App.propTypes = {};

ReactDom.render(<App />, document.getElementById("app"));
