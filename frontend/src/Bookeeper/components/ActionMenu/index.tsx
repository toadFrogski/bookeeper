import { Add, Notifications, MenuOpen as MenuOpenIcon, Menu as MenuIcon } from "@mui/icons-material";
import { Fab, ListItemIcon, ListItemText, ListSubheader, Menu, MenuItem, MenuList } from "@mui/material";
import { FC, useRef, useState } from "react";
import { useTranslation } from "react-i18next";

const ActionMenu: FC = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const menuRef = useRef();

  const [t] = useTranslation();

  return (
    <>
      <Fab
        ref={menuRef}
        color="primary"
        sx={{ position: "absolute", bottom: "24px", right: "24px" }}
        onClick={() => setIsMenuOpen(true)}
      >
        {isMenuOpen ? <MenuOpenIcon /> : <MenuIcon />}
      </Fab>
      <Menu
        anchorEl={menuRef.current}
        open={isMenuOpen}
        onClose={() => setIsMenuOpen(false)}
        anchorOrigin={{
          vertical: "top",
          horizontal: "left",
        }}
        transformOrigin={{
          vertical: "bottom",
          horizontal: "right",
        }}
      >
        <MenuList subheader={<ListSubheader>{t("common.actionMenu")}</ListSubheader>}>
          <MenuItem>
            <ListItemIcon>
              <Add />
            </ListItemIcon>
            <ListItemText>{t("common.addBook")}</ListItemText>
          </MenuItem>
          <MenuItem>
            <ListItemIcon>
              <Notifications />
            </ListItemIcon>
            <ListItemText>{t("common.seeNotifications")}</ListItemText>
          </MenuItem>
        </MenuList>
      </Menu>
    </>
  );
};

export default ActionMenu;
