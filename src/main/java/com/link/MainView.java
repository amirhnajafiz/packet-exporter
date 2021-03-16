package com.link;

import com.vaadin.flow.component.dependency.JavaScript;
import com.vaadin.flow.component.html.Div;
import com.vaadin.flow.router.Route;


@JavaScript("./frontend://script.js")
@Route
public class MainView extends Div {
    public MainView() {
        getElement().executeJs("meet($0)", "Client");
    }
}
