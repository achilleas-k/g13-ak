package keyboard

var (
	keysByName = map[string]int{
		"KeyEsc":              1,
		"Key1":                2,
		"Key2":                3,
		"Key3":                4,
		"Key4":                5,
		"Key5":                6,
		"Key6":                7,
		"Key7":                8,
		"Key8":                9,
		"Key9":                10,
		"Key0":                11,
		"KeyMinus":            12,
		"KeyEqual":            13,
		"KeyBackspace":        14,
		"KeyTab":              15,
		"KeyQ":                16,
		"KeyW":                17,
		"KeyE":                18,
		"KeyR":                19,
		"KeyT":                20,
		"KeyY":                21,
		"KeyU":                22,
		"KeyI":                23,
		"KeyO":                24,
		"KeyP":                25,
		"KeyLeftbrace":        26,
		"KeyRightbrace":       27,
		"KeyEnter":            28,
		"KeyLeftctrl":         29,
		"KeyA":                30,
		"KeyS":                31,
		"KeyD":                32,
		"KeyF":                33,
		"KeyG":                34,
		"KeyH":                35,
		"KeyJ":                36,
		"KeyK":                37,
		"KeyL":                38,
		"KeySemicolon":        39,
		"KeyApostrophe":       40,
		"KeyGrave":            41,
		"KeyLeftshift":        42,
		"KeyBackslash":        43,
		"KeyZ":                44,
		"KeyX":                45,
		"KeyC":                46,
		"KeyV":                47,
		"KeyB":                48,
		"KeyN":                49,
		"KeyM":                50,
		"KeyComma":            51,
		"KeyDot":              52,
		"KeySlash":            53,
		"KeyRightshift":       54,
		"KeyKpasterisk":       55,
		"KeyLeftalt":          56,
		"KeySpace":            57,
		"KeyCapslock":         58,
		"KeyF1":               59,
		"KeyF2":               60,
		"KeyF3":               61,
		"KeyF4":               62,
		"KeyF5":               63,
		"KeyF6":               64,
		"KeyF7":               65,
		"KeyF8":               66,
		"KeyF9":               67,
		"KeyF10":              68,
		"KeyNumlock":          69,
		"KeyScrolllock":       70,
		"KeyKp7":              71,
		"KeyKp8":              72,
		"KeyKp9":              73,
		"KeyKpminus":          74,
		"KeyKp4":              75,
		"KeyKp5":              76,
		"KeyKp6":              77,
		"KeyKpplus":           78,
		"KeyKp1":              79,
		"KeyKp2":              80,
		"KeyKp3":              81,
		"KeyKp0":              82,
		"KeyKpdot":            83,
		"KeyZenkakuhankaku":   85,
		"Key102Nd":            86,
		"KeyF11":              87,
		"KeyF12":              88,
		"KeyRo":               89,
		"KeyKatakana":         90,
		"KeyHiragana":         91,
		"KeyHenkan":           92,
		"KeyKatakanahiragana": 93,
		"KeyMuhenkan":         94,
		"KeyKpjpcomma":        95,
		"KeyKpenter":          96,
		"KeyRightctrl":        97,
		"KeyKpslash":          98,
		"KeySysrq":            99,
		"KeyRightalt":         100,
		"KeyLinefeed":         101,
		"KeyHome":             102,
		"KeyUp":               103,
		"KeyPageup":           104,
		"KeyLeft":             105,
		"KeyRight":            106,
		"KeyEnd":              107,
		"KeyDown":             108,
		"KeyPagedown":         109,
		"KeyInsert":           110,
		"KeyDelete":           111,
		"KeyMacro":            112,
		"KeyMute":             113,
		"KeyVolumedown":       114,
		"KeyVolumeup":         115,
		"KeyPower":            116, /*ScSystemPowerDown*/
		"KeyKpequal":          117,
		"KeyKpplusminus":      118,
		"KeyPause":            119,
		"KeyScale":            120, /*AlCompizScale(Expose)*/
		"KeyKpcomma":          121,
		"KeyHangeul":          122,
		"KeyHanja":            123,
		"KeyYen":              124,
		"KeyLeftmeta":         125,
		"KeyRightmeta":        126,
		"KeyCompose":          127,
		"KeyStop":             128, /*AcStop*/
		"KeyAgain":            129,
		"KeyProps":            130, /*AcProperties*/
		"KeyUndo":             131, /*AcUndo*/
		"KeyFront":            132,
		"KeyCopy":             133, /*AcCopy*/
		"KeyOpen":             134, /*AcOpen*/
		"KeyPaste":            135, /*AcPaste*/
		"KeyFind":             136, /*AcSearch*/
		"KeyCut":              137, /*AcCut*/
		"KeyHelp":             138, /*AlIntegratedHelpCenter*/
		"KeyMenu":             139, /*Menu(ShowMenu)*/
		"KeyCalc":             140, /*AlCalculator*/
		"KeySetup":            141,
		"KeySleep":            142, /*ScSystemSleep*/
		"KeyWakeup":           143, /*SystemWakeUp*/
		"KeyFile":             144, /*AlLocalMachineBrowser*/
		"KeySendfile":         145,
		"KeyDeletefile":       146,
		"KeyXfer":             147,
		"KeyProg1":            148,
		"KeyProg2":            149,
		"KeyWww":              150, /*AlInternetBrowser*/
		"KeyMsdos":            151,
		"KeyCoffee":           152, /*AlTerminalLock/Screensaver*/
		"KeyDirection":        153,
		"KeyCyclewindows":     154,
		"KeyMail":             155,
		"KeyBookmarks":        156, /*AcBookmarks*/
		"KeyComputer":         157,
		"KeyBack":             158, /*AcBack*/
		"KeyForward":          159, /*AcForward*/
		"KeyClosecd":          160,
		"KeyEjectcd":          161,
		"KeyEjectclosecd":     162,
		"KeyNextsong":         163,
		"KeyPlaypause":        164,
		"KeyPrevioussong":     165,
		"KeyStopcd":           166,
		"KeyRecord":           167,
		"KeyRewind":           168,
		"KeyPhone":            169, /*MediaSelectTelephone*/
		"KeyIso":              170,
		"KeyConfig":           171, /*AlConsumerControlConfiguration*/
		"KeyHomepage":         172, /*AcHome*/
		"KeyRefresh":          173, /*AcRefresh*/
		"KeyExit":             174, /*AcExit*/
		"KeyMove":             175,
		"KeyEdit":             176,
		"KeyScrollup":         177,
		"KeyScrolldown":       178,
		"KeyKpleftparen":      179,
		"KeyKprightparen":     180,
		"KeyNew":              181, /*AcNew*/
		"KeyRedo":             182, /*AcRedo/Repeat*/
		"KeyF13":              183,
		"KeyF14":              184,
		"KeyF15":              185,
		"KeyF16":              186,
		"KeyF17":              187,
		"KeyF18":              188,
		"KeyF19":              189,
		"KeyF20":              190,
		"KeyF21":              191,
		"KeyF22":              192,
		"KeyF23":              193,
		"KeyF24":              194,
		"KeyPlaycd":           200,
		"KeyPausecd":          201,
		"KeyProg3":            202,
		"KeyProg4":            203,
		"KeyDashboard":        204, /*AlDashboard*/
		"KeySuspend":          205,
		"KeyClose":            206, /*AcClose*/
		"KeyPlay":             207,
		"KeyFastforward":      208,
		"KeyBassboost":        209,
		"KeyPrint":            210, /*AcPrint*/
		"KeyHp":               211,
		"KeyCamera":           212,
		"KeySound":            213,
		"KeyQuestion":         214,
		"KeyEmail":            215,
		"KeyChat":             216,
		"KeySearch":           217,
		"KeyConnect":          218,
		"KeyFinance":          219, /*AlCheckbook/Finance*/
		"KeySport":            220,
		"KeyShop":             221,
		"KeyAlterase":         222,
		"KeyCancel":           223, /*AcCancel*/
		"KeyBrightnessdown":   224,
		"KeyBrightnessup":     225,
		"KeyMedia":            226,
		"KeySwitchvideomode":  227, /*CycleBetweenAvailableVideo */
		"KeyKbdillumtoggle":   228,
		"KeyKbdillumdown":     229,
		"KeyKbdillumup":       230,
		"KeySend":             231, /*AcSend*/
		"KeyReply":            232, /*AcReply*/
		"KeyForwardmail":      233, /*AcForwardMsg*/
		"KeySave":             234, /*AcSave*/
		"KeyDocuments":        235,
		"KeyBattery":          236,
		"KeyBluetooth":        237,
		"KeyWlan":             238,
		"KeyUwb":              239,
		"KeyUnknown":          240,
		"KeyVideoNext":        241, /*DriveNextVideoSource*/
		"KeyVideoPrev":        242, /*DrivePreviousVideoSource*/
		"KeyBrightnessCycle":  243, /*BrightnessUp,AfterMaxIsMin*/
		"KeyBrightnessZero":   244, /*BrightnessOff,UseAmbient*/
		"KeyDisplayOff":       245, /*DisplayDeviceToOffState*/
		"KeyWimax":            246,
		"KeyRfkill":           247, /*KeyThatControlsAllRadios*/
		"KeyMicmute":          248, /*Mute/UnmuteTheMicrophone*/
	}
)

func KeyCode(name string) int {
	return keysByName[name]
}