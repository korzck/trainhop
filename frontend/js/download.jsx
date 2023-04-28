class HelloWorld extends React.Component {
    render() {
        return (
            <div className='container'>
                <div className='row'>
                    <div className='col-xs-8 col-xs-offset-2 jumbotron text-center'>
                        <h3>trainhop prediction model</h3>
                        <h3>coded/tested/deployed by PO-sqrt-team</h3>
                        <h3>resulting file</h3>

                        <h4>
                            <a href='/download'>link to download</a>
                            <br />
                            <a href='https://github.com/korzck/trainhop'>
                                Github
                            </a>
                        </h4>
                    </div>
                </div>
            </div>
        );
    }
}

ReactDOM.render(<HelloWorld />, document.getElementById("download"));
