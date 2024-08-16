package main

import "log"

func DispatchEvent(event StreamDeckEvent) {
	log.Printf("Dispatching event: %+v\n", event)

	if event.IsActionAssociated() {
		if actionEvent, ok := event.(interface {
			GetAction() (string, bool)
		}); ok {
			actionUUID, actionExists := actionEvent.GetAction()
			if !actionExists {
				log.Printf("No action registered for UUID: %s", actionUUID)
				return
			}
			action, exists := actionRegistry[actionUUID]
			if !exists {
				log.Printf("No action registered for UUID: %s", actionUUID)
				return
			}

			// Dynamically call the appropriate handler if it exists
			switch e := event.(type) {
			case *DidReceiveSettingsEvent:
				if handler, ok := action.(interface {
					HandleDidReceiveSettings(*DidReceiveSettingsEvent)
				}); ok {
					handler.HandleDidReceiveSettings(e)
				}
			case *KeyDownEvent:
				log.Print("KeyDownEvent")
				if handler, ok := action.(interface {
					HandleKeyDown(*KeyDownEvent)
				}); ok {
					handler.HandleKeyDown(e)
				}

			case *KeyUpEvent:
				if handler, ok := action.(interface {
					HandleKeyUp(*KeyUpEvent)
				}); ok {
					handler.HandleKeyUp(e)
				}
			case *WillAppearEvent:
				if handler, ok := action.(interface {
					HandleWillAppear(*WillAppearEvent)
				}); ok {
					handler.HandleWillAppear(e)
				}
			case *WillDisappearEvent:
				if handler, ok := action.(interface {
					HandleWillDisappear(*WillDisappearEvent)
				}); ok {
					handler.HandleWillDisappear(e)
				}
			case *TitleParametersDidChangeEvent:
				if handler, ok := action.(interface {
					HandleTitleParametersDidChange(*TitleParametersDidChangeEvent)
				}); ok {
					handler.HandleTitleParametersDidChange(e)
				}
			case *TouchTapEvent:
				if handler, ok := action.(interface {
					HandleTouchTap(*TouchTapEvent)
				}); ok {
					handler.HandleTouchTap(e)
				}
			case *DialDownEvent:
				if handler, ok := action.(interface {
					HandleDialDown(*DialDownEvent)
				}); ok {
					handler.HandleDialDown(e)
				}
			case *DialUpEvent:
				if handler, ok := action.(interface {
					HandleDialUp(*DialUpEvent)
				}); ok {
					handler.HandleDialUp(e)
				}
			case *DialRotateEvent:
				if handler, ok := action.(interface {
					HandleDialRotate(*DialRotateEvent)
				}); ok {
					handler.HandleDialRotate(e)
				}
			case *PropertyInspectorDidAppearEvent:
				if handler, ok := action.(interface {
					HandlePropertyInspectorDidAppear(*PropertyInspectorDidAppearEvent)
				}); ok {
					handler.HandlePropertyInspectorDidAppear(e)
				}
			case *PropertyInspectorDidDisappearEvent:
				if handler, ok := action.(interface {
					HandlePropertyInspectorDidDisappear(*PropertyInspectorDidDisappearEvent)
				}); ok {
					handler.HandlePropertyInspectorDidDisappear(e)
				}
			case *SendToPluginEvent:
				if handler, ok := action.(interface {
					HandleSendToPlugin(*SendToPluginEvent)
				}); ok {
					handler.HandleSendToPlugin(e)
				}
			case *SendToPropertyInspectorEvent:
				if handler, ok := action.(interface {
					HandleSendToPropertyInspector(*SendToPropertyInspectorEvent)
				}); ok {
					handler.HandleSendToPropertyInspector(e)
				}
			default:
				log.Printf("No handler found for action-associated event type: %T", e)
			}
		} else {
			log.Printf("Failed to cast event to ActionAssociatedEvent type")
		}
	} else {
		// Handle global event
		for _, action := range actionRegistry {
			switch e := event.(type) {
			case *DidReceiveGlobalSettingsEvent:
				if handler, ok := action.(interface {
					HandleDidReceiveGlobalSettings(*DidReceiveGlobalSettingsEvent)
				}); ok {
					handler.HandleDidReceiveGlobalSettings(e)
				}
			case *DidReceiveDeepLinkEvent:
				if handler, ok := action.(interface {
					HandleDidReceiveDeepLink(*DidReceiveDeepLinkEvent)
				}); ok {
					handler.HandleDidReceiveDeepLink(e)
				}
			case *DeviceDidConnectEvent:
				if handler, ok := action.(interface {
					HandleDeviceDidConnect(*DeviceDidConnectEvent)
				}); ok {
					handler.HandleDeviceDidConnect(e)
				}
			case *DeviceDidDisconnectEvent:
				if handler, ok := action.(interface {
					HandleDeviceDidDisconnect(*DeviceDidDisconnectEvent)
				}); ok {
					handler.HandleDeviceDidDisconnect(e)
				}
			case *ApplicationDidLaunchEvent:
				if handler, ok := action.(interface {
					HandleApplicationDidLaunch(*ApplicationDidLaunchEvent)
				}); ok {
					handler.HandleApplicationDidLaunch(e)
				}
			case *ApplicationDidTerminateEvent:
				if handler, ok := action.(interface {
					HandleApplicationDidTerminate(*ApplicationDidTerminateEvent)
				}); ok {
					handler.HandleApplicationDidTerminate(e)
				}
			case *SystemDidWakeUpEvent:
				if handler, ok := action.(interface {
					HandleSystemDidWakeUp(*SystemDidWakeUpEvent)
				}); ok {
					handler.HandleSystemDidWakeUp(e)
				}
			default:
				log.Printf("No handler found for global event type: %T", e)
			}
		}
	}
}
