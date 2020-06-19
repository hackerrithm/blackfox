import * as React from "react";
import { render } from "react-dom";
import App from "./app";
import ThemeProvider from "./components/theme/themeprovider";

render(
	<ThemeProvider>
		<App />
	</ThemeProvider>,
	document.getElementById("root")
);
