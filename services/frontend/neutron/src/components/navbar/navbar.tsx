import React, { useReducer } from "react";

import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import InputBase from "@material-ui/core/InputBase";
import Badge from "@material-ui/core/Badge";
import MenuItem from "@material-ui/core/MenuItem";
import Menu from "@material-ui/core/Menu";
import MenuIcon from "@material-ui/icons/Menu";
import SearchIcon from "@material-ui/icons/Search";
import AccountCircle from "@material-ui/icons/AccountCircle";
import MailIcon from "@material-ui/icons/Mail";
import NotificationsIcon from "@material-ui/icons/Notifications";
import MoreIcon from "@material-ui/icons/MoreVert";
import { getThemeProps } from "@material-ui/styles";
import { Link, useHistory } from "react-router-dom";
import "./style.css";
import { authContext } from "../utils/authContext";
import { LoginReducer, InitialState } from "../login";
import Button from "@material-ui/core/Button";
import useStyles from "./styles";

export const Context = React.createContext(InitialState);

const PrimarySearchAppBar: React.FC = () => {
	const { auth } = React.useContext(authContext);
	const [state, dispatch] = useReducer(LoginReducer, InitialState);
	const history = useHistory();
	const classes = useStyles(getThemeProps);
	const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
	const [
		mobileMoreAnchorEl,
		setMobileMoreAnchorEl
	] = React.useState<null | HTMLElement>(null);

	const isMenuOpen = Boolean(anchorEl);
	const isMobileMenuOpen = Boolean(mobileMoreAnchorEl);

	const handleProfileMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
		setAnchorEl(event.currentTarget);
	};

	const handleMobileMenuClose = () => {
		setMobileMoreAnchorEl(null);
	};

	const handleMenuClose = () => {
		setAnchorEl(null);
		handleMobileMenuClose();
	};

	const handleMobileMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
		setMobileMoreAnchorEl(event.currentTarget);
	};

	const menuId = "primary-search-account-menu";
	const renderMenu = (
		<Menu
			anchorEl={anchorEl}
			anchorOrigin={{ vertical: "top", horizontal: "right" }}
			id={menuId}
			keepMounted
			transformOrigin={{ vertical: "top", horizontal: "right" }}
			open={isMenuOpen}
			onClose={handleMenuClose}
		>
			<MenuItem onClick={handleMenuClose}>Profile</MenuItem>
			<MenuItem onClick={handleMenuClose}>My account</MenuItem>
		</Menu>
	);

	const mobileMenuId = "primary-search-account-menu-mobile";
	const renderMobileMenu = (
		<Menu
			anchorEl={mobileMoreAnchorEl}
			anchorOrigin={{ vertical: "top", horizontal: "right" }}
			id={mobileMenuId}
			keepMounted
			transformOrigin={{ vertical: "top", horizontal: "right" }}
			open={isMobileMenuOpen}
			onClose={handleMobileMenuClose}
		>
			<MenuItem>
				<IconButton aria-label="show 4 new mails" color="inherit">
					<Badge badgeContent={4} color="secondary">
						<MailIcon />
					</Badge>
				</IconButton>
				<p>Messages</p>
			</MenuItem>
			<MenuItem>
				<Link to="/about">
					<IconButton
						aria-label="show 11 new notifications"
						color="inherit"
					>
						<Badge badgeContent={11} color="secondary">
							<NotificationsIcon />
						</Badge>
					</IconButton>
				</Link>
				<p>Notifications</p>
			</MenuItem>
			<MenuItem onClick={handleProfileMenuOpen}>
				<IconButton
					aria-label="account of current user"
					aria-controls="primary-search-account-menu"
					aria-haspopup="true"
					color="inherit"
				>
					<AccountCircle />
				</IconButton>
				<p>Profile</p>
			</MenuItem>
		</Menu>
	);

	return (
		<Context.Provider
			value={{
				username: state.username,
				isLoading: state.isLoading,
				isLoggedIn: state.isLoggedIn,
				error: state.error
			}}
		>
			<div className={classes.grow}>
				<AppBar
					color="secondary"
					style={{ backgroundColor: "white" }}
					position="sticky"
				>
					<Toolbar variant="dense">
						<IconButton
							edge="start"
							className={classes.menuButton}
							color="secondary"
							aria-label="open drawer"
						>
							<MenuIcon />
						</IconButton>
						<Link to="/" style={{ textDecoration: "none" }}>
							<img
								style={{ width: "50px" }}
								src="https://icon-library.net/images/fox-icon-png/fox-icon-png-12.jpg"
								alt="Blackfox"
							/>
						</Link>
						<Link
									to="/discover"
									style={{
										textDecoration: "none",
										marginLeft: 10,
										color: "black"
									}}
								>
									<Button color="inherit">discover</Button>
								</Link>
								<Link
									to="/browse"
									style={{
										textDecoration: "none",
										marginLeft: 10,
										color: "black"
									}}
								>
									<Button color="inherit">browse</Button>
								</Link>
						<div className={classes.search}>
							<div className={classes.searchIcon}>
								<SearchIcon />
							</div>
							<InputBase
								placeholder="Search…"
								classes={{
									root: classes.inputRoot,
									input: classes.inputInput
								}}
								inputProps={{ "aria-label": "search" }}
							/>
						</div>
						<div className={classes.grow} />
						<div className={classes.sectionDesktop}>
							{state.isLogged ? (
								<IconButton
									aria-label="show 4 new mails"
									color="secondary"
								>
									<Badge badgeContent={4} color="secondary">
										<MailIcon />
									</Badge>
								</IconButton>
							) : null}
							{!state.isLoggedIn && (
								<Link
									to="/about"
									style={{
										textDecoration: "none",
										marginLeft: 10,
										color: "black"
									}}
								>
									<Button color="inherit">about</Button>
								</Link>
							)}
							<Link
								to="/register"
								style={{ textDecoration: "none" }}
							>
								<Button
									color="inherit"
									variant="outlined"
									style={{ marginLeft: 10, color: "black" }}
								>
									register
								</Button>
							</Link>
							<Link
								to="/login"
								style={{ textDecoration: "none" }}
							>
								<Button
									color="primary"
									variant="contained"
									style={{
										marginLeft: 10,
										backgroundColor: "#5588FF"
									}}
								>
									login
								</Button>
							</Link>
							<Link
								to="/"
								style={{
									textDecoration: "none",
									marginLeft: 10,
									color: "black"
								}}
								onClick={() => {
									localStorage.removeItem("userAuthData");
									dispatch({ type: "logOut" });
									history.push("/");
								}}
							>
								<Button color="inherit">logout</Button>
							</Link>
						</div>
						<div className={classes.sectionMobile}>
							<IconButton
								aria-label="show more"
								aria-controls={mobileMenuId}
								aria-haspopup="true"
								onClick={handleMobileMenuOpen}
								color="secondary"
							>
								<MoreIcon />
							</IconButton>
						</div>
					</Toolbar>
				</AppBar>
				{renderMobileMenu}
				{renderMenu}
			</div>
		</Context.Provider>
	);
};

export default PrimarySearchAppBar;
