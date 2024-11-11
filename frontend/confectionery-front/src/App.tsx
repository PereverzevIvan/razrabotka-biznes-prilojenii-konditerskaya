import Container from "./components/Container/Container";
import Footer from "./modules/Footer/Footer";
import Header from "./modules/Header/Header";
import "./scss/index.scss";

function App() {
  return (
    <>
      <div className="App">
        <Header />
        <main className="main">
          <Container>Привет, это проверка</Container>
        </main>
        <Footer />
      </div>
    </>
  );
}

export default App;
