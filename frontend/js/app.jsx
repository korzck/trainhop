class HelloWorld extends React.Component {
    render() {
        return (
            <div className='container'>
                <div className='row'>
                    <div className='col-xs-8 col-xs-offset-2 jumbotron text-center'>
                        <h3>trainhop prediction model</h3>
                        <h3>Coded/tested/deployed by PO-sqrt-team</h3>
                        <form
                            action='/upload'
                            method='post'
                            enctype='multipart/form-data'
                        >
                            <input type='file' name='file' />
                            <input type='submit' />
                        </form>

                        <h4>
                            <a href='https://github.com/korzck'>Github</a>
                        </h4>
                    </div>
                </div>
            </div>
        );
    }
}

ReactDOM.render(<HelloWorld />, document.getElementById("app"));
