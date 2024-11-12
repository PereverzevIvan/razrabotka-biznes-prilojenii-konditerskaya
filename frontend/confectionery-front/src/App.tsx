import Container from "./components/Container/Container";
import Footer from "./modules/Footer/Footer";
import Header from "./modules/Header/Header";
import CustomRouter from "./modules/Router/Router";
import "./scss/index.scss";

function App() {
  return (
    <>
      <div className="App">
        <Header />
        <main className="main">
          <Container>
            <CustomRouter />
          </Container>
        </main>
        <Footer />
      </div>
    </>
  );
}

export default App;
