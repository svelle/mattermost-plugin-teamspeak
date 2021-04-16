export interface PluginRegistry {
    registerPostTypeComponent(typeName: string, component: React.ElementType)

    // Add more if needed from https://developers.mattermost.com/extend/plugins/webapp/reference

    /**
    * Register a Right-Hand Sidebar component by providing a title for the right hand component.
    * Accepts the following:
    * - title - A string or JSX element to display as a title for the RHS.
    * - component - A react component to display in the Right-Hand Sidebar.
    * Returns:
    * - id: a unique identifier
    * - showRHSPlugin: the action to dispatch that will open the RHS.
    * - hideRHSPlugin: the action to dispatch that will close the RHS
    * - toggleRHSPlugin: the action to dispatch that will toggle the RHS
    */
    registerRightHandSidebarComponent(component, title)

    /**
    * Add a button to the channel header. If there are more than one buttons registered by any
    * plugin, a dropdown menu is created to contain all the plugin buttons.
    * Accepts the following:
    * - icon - React element to use as the button's icon
    * - action - a function called when the button is clicked, passed the channel and channel member as arguments
    * - dropdown_text - string or React element shown for the dropdown button description
    * - tooltip_text - string shown for tooltip appear on hover
    */
    registerChannelHeaderButtonAction(icon, action, dropdownText, tooltipText)
}
