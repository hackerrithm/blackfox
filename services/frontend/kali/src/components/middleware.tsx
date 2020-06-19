import React, { useState } from "react";
import { MuiThemeProvider, createMuiTheme } from "@material-ui/core/styles";
import App from "../app";

const Middleware = () => {
	const [theme, setTheme] = useState<any>({
		palette: {
                        type: "dark",
		}
	});

	const toggleDarkTheme = () => {
		let newPaletteType = theme.palette.type === "light" ? "dark" : "light";
		setTheme({
			palette: {
				type: newPaletteType
			}
		});
	};

	const muiTheme = createMuiTheme(theme);

	return (
		<MuiThemeProvider theme={muiTheme}>
			<App onToggleDark={toggleDarkTheme} />
		</MuiThemeProvider>
	);
};

export default Middleware;
