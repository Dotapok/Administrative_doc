package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// structures
type ActeNaissance struct {
	Numero         string `json:"Nº"`
	Type           string `json:"Type de document"`
	PROVINCE       string `json:"PROVINCE"`
	DEPARTMENT     string `json:"DEPARTEMENT"`
	ARRONDISSEMENT string `json:"ARRONDISSEMENT"`
	De             string `json:"De"`

	NomEnfant    string `json:"Nom de l’enfant"`
	PrenomEnfant string `json:"Prénom(s) de l’enfant"`
	EstNeA       string `json:"Lieu de naissance de l'enfant"`
	Le           string `json:"Date de naissance de l'enfant"`
	Sexe         string `json:"Sexe de l'enfant"`

	De0        string `json:"Nom du père"`
	NeLe0      string `json:"Date de naissance du père"`
	NeA        string `json:"Lieu de naissance du père"`
	Natio0     string `json:"Nationalité du père"`
	DomicilieA string `json:"Domicile du père"`
	Profession string `json:"Profession du père"`

	De1         string `json:"Nom de la mère"`
	NeLe1       string `json:"Date de naissance de la mère"`
	NeeA        string `json:"Lieu de naissance de la mère"`
	Natio1      string `json:"Nationalité de la mère"`
	DomicilieeA string `json:"Domicile de la mère"`
	Profession2 string `json:"Profession de la mère"`

	DeclarationDe     string    `json:"Sur la declaration de"`
	ParNous           string    `json:"Par nous"`
	Centre            string    `json:"Officier De l’état civil du centre de"`
	AssisteDe         string    `json:"Assisté de"`
	SignatureOfficier string    `json:"Signature de l'officier"`
	NumeroDEC         string    `json:"Numero de la declaration initiatrice "`
	DresseLe          time.Time `json:"Dressé le"`
}

type ActeDeces struct {
	Numero         string `json:"Nº"`
	Type           string `json:"Type"`
	PROVINCE       string `json:"PROVINCE"`
	DEPARTMENT     string `json:"DEPARTEMENT"`
	ARRONDISSEMENT string `json:"ARRONDISSEMENT"`
	De             string `json:"De"`

	NomDefunt       string `json:"Nom du defunt/de la defunte"`
	PrenomDefunt    string `json:"Prénom(s) du defunt/de la defunte"`
	EstDecesA       string `json:"Lieu de survenu du deces du defunt/de la defunte"`
	DecedeLe        string `json:"Date de survenue du decès du defunt/de la defunte"`
	EstNeA          string `json:"Lieu de naissance du defunt/de la defunte"`
	Le              string `json:"Date de naissance du defunt/de la defunte"`
	Sexe            string `json:"Sexe du defunt/de la defunte"`
	Natio           string `json:"Nationalité du defunt/de la defunte"`
	Profession      string `json:"Profession du defunt/de la defunte"`
	DomicilieDefunt string `json:"Domicile du defunt/defunte"`

	NP         string `json:"Nom du père du defunt/de la defunte"`
	NM         string `json:"Nom de la mère du defunt/de la defunte"`
	DomicilieA string `json:"Domicile des parents"`

	DeclarationDe     string    `json:"Sur la declaration de"`
	ParNous           string    `json:"Par nous"`
	Centre            string    `json:"Officier De l’état civil du centre de"`
	AssisteDe         string    `json:"Assisté de"`
	SignatureOfficier string    `json:"Signature de l'officier"`
	NumeroDEC         string    `json:"Numero de la declaration initiatrice "`
	DresseLe          time.Time `json:"Dressé le"`
}

type ActeMariage struct {
	Numero         string `json:"Nº"`
	Type           string `json:"Type"`
	PROVINCE       string `json:"PROVINCE"`
	DEPARTMENT     string `json:"DEPARTEMENT"`
	ARRONDISSEMENT string `json:"ARRONDISSEMENT"`
	De             string `json:"De"`

	Contracter string `json:"Contracter le"`

	NomEpoux     string `json:"Nom(s) de l'epoux"`
	PrenomEpoux  string `json:"Prénom(s) de l'epoux'"`
	NomPereH     string `json:"Nom du père de l'epoux"`
	NomMereH     string `json:"Nom de la mère de l'epoux"`
	DateNEpoux   string `json:"Date de naissance de l'epoux"`
	LieuNEpoux   string `json:"Lieu de naissance de l'epoux"`
	RaceH        string `json:"Race de l'epoux"`
	GroupemmentH string `json:"Groupement de l'epoux"`
	SubdivisionH string `json:"Subdivision de l'epoux"`
	RegionH      string `json:"Region de l'epoux"`
	ProffEPH     string `json:"Profession de l'epoux"`
	ResidenceH   string `json:"Residence de l'epoux"`

	NomEpouse    string `json:"Nom(s) de l'epouse"`
	PrenomEpouse string `json:"Prénom(s) de l'epouse'"`
	NomPereF     string `json:"Nom du père de l'epouse"`
	NomMereF     string `json:"Nom de la mère de l'epouse"`
	DateNEpouse  string `json:"Date de naissance de l'epouse"`
	LieuNEpouse  string `json:"Lieu de naissance de l'epouse"`
	RaceF        string `json:"Race de l'epouse"`
	GroupemmentF string `json:"Groupement de l'epouse"`
	SubdivisionF string `json:"Subdivision de l'epouse"`
	RegionF      string `json:"Region de l'epouse"`
	ProffEPF     string `json:"Profession de l'epouse"`
	ResidenceF   string `json:"Residence de l'epouse"`

	ConsentementEPOUX   string `json:"Consentement des epoux"`
	ConsentementCHEFFam string `json:"Consentement du chef de famille"`
	ConsentementEpoux   string `json:"Consentement de l'epoux"`
	ConsentementEpouse  string `json:"Consentement de l'epouse"`
	Oppositions         string `json:"Oppositions"`
	MontantDOT          string `json:"Montant convenu pour la dot"`
	SommeDOTVERSER      string `json:"Somme versée"`
	DateVerse           string `json:"Date de versement"`
	DateVerseCOM        string `json:"Date et montant des versements complementaires"`

	TemoinsH string `json:"Temoins du marié"`
	TemoinsF string `json:"Temoins de la marié"`

	ParNous           string    `json:"Par nous"`
	Centre            string    `json:"Officier De l’état civil du centre de"`
	AssisteDe         string    `json:"Assisté de"`
	SignatureOfficier string    `json:"Signature de l'officier"`
	NumeroDEC         string    `json:"Numero de la declaration initiatrice "`
	DresseLe          time.Time `json:"Dressé le"`
}

// liste des declarations

type DeclarationNaissance struct {
	Numero        string `json:"N°"`
	Type          string `json:"Type de document"`
	Etablissement string `json:"Designation de l'etablissement"`
	Le            string `json:"Date de declaration"`
	A             string `json:"Lieu de declaration"`
	Heure         string `json:"heures de declaration"`
	PrenomEnfant  string `json:"Prénom(s) de l’enfant"`
	Nom           string `json:"Nom de l'enfant"`
	Sexe          string `json:"Sexe de l'enfant"`

	De0        string `json:"Nom du pere de l'enfant"`
	Pre0       string `json:"Prenom du pere de l'enfant"`
	NeLe0      string `json:"Date de naissance du père de l'enfant"`
	NeA        string `json:"Lieu de naissance du père de l'enfant"`
	Profession string `json:"Profession du père de l'enfant"`
	CniP       string `json:"CNI du père de l'enfant"`

	De1         string `json:"Nom de la mere de l'enfant"`
	Pre1        string `json:"Prenom de la mere de l'enfant"`
	NeLe1       string `json:"Date de naissance de la mère de l'enfant"`
	NeeA        string `json:"Lieu de naissance de la mère de l'enfant"`
	Profession2 string `json:"Profession de la mère de l'enfant"`
	Dom         string `json:"Domiciles des parents "`
	SM          string `json:"Situation matrimoniale des parents"`
	CniM        string `json:"CNI de la mère de l'enfant"`

	Ville              string    `json:"Ville de la declaration"`
	Tem                string    `json:"Temoins "`
	DateDEC            string    `json:"Date de la declaration"`
	SignatureDeclarant string    `json:"Signature "`
	Signataire         string    `json:"Signataire de la declaration"`
	NumCNIDEC          string    `json:"Numero de CNI du declarant "`
	ProfDEC            string    `json:"Profession du declarant"`
	DateENREG          time.Time `json:"Date enregistrement definitif"`
}

type DeclarationDeces struct {
	Numero            string    `json:"N°"`
	Type              string    `json:"Type de document"`
	Nom_def           string    `json:"Nom du defunt/de la defunte"`
	Date_naissance    string    `json:"Date de naissance du defunt/de la defunte"`
	Lieu_naissance    string    `json:"Lieu de naissance du defunt/de la defunte"`
	Date_deces        string    `json:"Date de decès"`
	Heure_deces       string    `json:"Heure de decès"`
	Lieu_deces        string    `json:"Lieu de decès"`
	Cause_deces       string    `json:"Cause du decès"`
	Profession        string    `json:"Profession du defunt/de la defunte"`
	Nom_temoins       string    `json:"Temoins "`
	Ville_declaration string    `json:"Ville de la declaration "`
	Sexe              string    `json:"Sexe du defunt/de la defunte"`
	Nationalite       string    `json:"Nationalite du defunt/de la defunte"`
	Etat_civil        string    `json:"Etat civil du defunt/ de la defunte"`
	Numero_identite   string    `json:"Numero CNI defunt/defunte"`
	DateELAB          string    `json:"Date d'elaboration de la declaration"`
	HeureELAB         string    `json:"Heure d'elaboration de la declaration"`
	NumCNIDEC         string    `json:"Numero CNI declarant "`
	ProfDEC           string    `json:"Profession du declarant"`
	Signataire        string    `json:"Nom du declarant "`
	Signature         string    `json:"Signature"`
	DateENREG         time.Time `json:"Date enregistrement definitif"`
}

type PublicationBans struct {
	Numero           string `json:"N°"`
	Type             string `json:"Type de document"`
	Date_debut       string `json:"Date debut de la publication"`
	Date_fin         string `json:"Date fin de la publication"`
	Lieu_publication string `json:"Lieu de publication"`

	Nom_epoux             string `json:"Nom de l'epoux"`
	Prenoms_epoux         string `json:"Prénom(s) de l'epoux"`
	Date_naissance_epoux  string `json:"Date de naissance epoux"`
	Lieu_naissance_epoux  string `json:"Lieu de naissance de l'epoux"`
	Profession_epoux      string `json:"Profession de l'epoux "`
	Nationalite_epoux     string `json:"Nationalité de l'epoux"`
	Domicile_epoux        string `json:"Domicile de l'epoux"`
	EtatCivil_epoux       string `json:"Etat civil de l'epoux"`
	Numero_identite_epoux string `json:"Numero CNI epoux"`

	Nom_epouse             string `json:"Nom de l'epouse"`
	Prenoms_epouse         string `json:"Prénom(s) de l'epouse"`
	Date_naissance_epouse  string `json:"Date de naissance epouse"`
	Lieu_naissance_epouse  string `json:"Lieu de naissance de l'epouse"`
	Profession_epouse      string `json:"Profession de l'epouse "`
	Nationalite_epouse     string `json:"Nationalité de l'epouse"`
	Domicile_epouse        string `json:"Domicile de l'epouse"`
	EtatCivil_epouse       string `json:"Etat civil de l'epouse"`
	Numero_identite_epouse string `json:"Numero CNI epouse"`

	Date_mariage       string `json:"Date de mariage"`
	HeureDebut_mariage string `json:"Heure de debut du mariage"`
	DetailsMariage     string `json:"Details sur la ceremonie"`
	NomPereMarieHomme  string `json:"Nom du père du mari"`
	NomMereMarieHomme  string `json:"Nom de la mère du mari "`
	NomPereMarieFemme  string `json:"Nom du père de la marié"`
	NomMereMarieFemme  string `json:"Nom de la mère de la marié"`
	MunicipalitePub    string `json:"Municipalité de la publication"`

	Signataire        string    `json:"Signataire"`
	Signature         string    `json:"Signature"`
	DateELAB          string    `json:"Date d'elaboration"`
	ContactMunicipale string    `json:"Contact en cas de problème"`
	DateENREG         time.Time `json:"Date enregistrement definitif"`
}

// creation de la structure du contrat
type Administrative_docChaincode struct{}

func (t *Administrative_docChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *Administrative_docChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "creerActeNaissance" {
		return t.creerActeNaissance(stub, args)
	} else if function == "RecupActeNaissance" {
		return t.RecupActeNaissance(stub, args)
	} else if function == "ConsulterActesNaissance" {
		return t.consulterActesNaissance(stub)
	} else if function == "creerDeclarationNaissance" {
		return t.creerdeclaration(stub, args)
	} else if function == "RecupereDeclarationNaissance" {
		return t.Recupdeclaration(stub, args)
	} else if function == "ConsulterDeclarationNaissance" {
		return t.consulterDeclarationsNAISS(stub)
	} else if function == "creerDeclarationsDeces" {
		return t.creerDeclarationsDeces(stub, args)
	} else if function == "RecupererDeclarationDeces" {
		return t.RecupDeclarationDeces(stub, args)
	} else if function == "consulterDeclarationsDeces" {
		return t.consulterDeclarationsDeces(stub)
	} else if function == "creerActeDeces" {
		return t.creerActeDeces(stub, args)
	} else if function == "RecupActeDeces" {
		return t.RecupActeDeces(stub, args)
	} else if function == "consulterActesDeces" {
		return t.consulterActesDeces(stub)
	} else if function == "EnregistrerPublicationBans" {
		return t.EnregistrerPUB(stub, args)
	} else if function == "RecupererPublicationBans" {
		return t.RecupererPublicationBans(stub, args)
	} else if function == "consulterPUB" {
		return t.ConsulterPubBans(stub)
	} else if function == "creerActeMariage" {
		return t.CreerActesMariage(stub, args)
	} else if function == "RecupMariage" {
		return t.RecupActeMariage(stub, args)
	} else if function == "ConsulterMariage" {
		return t.consulterActesMariage(stub)
	}
	return shim.Error("Nom de methode appelée invalide.")
}

// acte de naissance
func (t *Administrative_docChaincode) creerActeNaissance(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 28 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	acteNaissance := ActeNaissance{
		Numero:         args[0],
		Type:           "acte de naissance",
		PROVINCE:       args[1],
		DEPARTMENT:     args[2],
		ARRONDISSEMENT: args[3],
		De:             args[4],

		NomEnfant:    args[5],
		PrenomEnfant: args[6],
		EstNeA:       args[7],
		Le:           args[8],
		Sexe:         args[9],

		De0:        args[10],
		NeLe0:      args[11],
		NeA:        args[12],
		Natio0:     args[13],
		DomicilieA: args[14],
		Profession: args[15],

		De1:         args[16],
		NeLe1:       args[17],
		NeeA:        args[18],
		Natio1:      args[19],
		DomicilieeA: args[20],
		Profession2: args[21],

		DeclarationDe:     args[22],
		ParNous:           args[23],
		Centre:            args[24],
		AssisteDe:         args[25],
		SignatureOfficier: args[26],
		NumeroDEC:         args[27],
		DresseLe:          time.Now(),
	}

	actesNaissanceAsBytes, err := stub.GetState(acteNaissance.Numero)
	if actesNaissanceAsBytes != nil {
		return shim.Error("Acte de naissance avec le numero suivant deja existant : " + acteNaissance.Numero)
	}

	Versionbyte, err := json.Marshal(acteNaissance)
	if err != nil {
		return shim.Error("Impossible de monter la structure de acte de naissance : " + err.Error())
	}

	err = stub.PutState(acteNaissance.Numero, Versionbyte)
	if err != nil {
		return shim.Error("Impossible d'enregistrer l'acte de naissance: " + err.Error())
	}
	return shim.Success(nil)
}

func (t *Administrative_docChaincode) RecupActeNaissance(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	acteNaissanceAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Impossible de trouver cette acte de naissance: " + err.Error())
	} else if acteNaissanceAsBytes == nil {
		return shim.Error("Acte de naissance: " + args[0] + " inexistant.")
	}

	return shim.Success(acteNaissanceAsBytes)
}

func (t *Administrative_docChaincode) consulterActesNaissance(stub shim.ChaincodeStubInterface) peer.Response {
	resultiterator, err := stub.GetQueryResult(`{"selector":{"Type":"acte de naissance"}}`)

	if err != nil {
		return shim.Error("Erreur lors de la recuperation des données" + err.Error())
	}
	defer resultiterator.Close()

	var results []ActeNaissance

	for resultiterator.HasNext() {
		queryReponse, err := resultiterator.Next()
		if err != nil {
			return shim.Error("Erreur lors de la recuperation des données" + err.Error())
		}
		var ficheActe ActeNaissance
		err = json.Unmarshal(queryReponse.Value, &ficheActe)
		if err != nil {
			return shim.Error("Erreur lors de la recuperation des données")
		}

		results = append(results, ficheActe)
	}

	resultatsAsbytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error("Erreur lors de la conversion en JSON des actes de naissance")
	}
	return shim.Success(resultatsAsbytes)
}

// declaration de naissance
func (t *Administrative_docChaincode) creerdeclaration(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 29 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	declarationN := DeclarationNaissance{
		Type:               "declaration de naissance",
		Etablissement:      args[0],
		Le:                 args[1],
		A:                  args[2],
		Heure:              args[3],
		Numero:             args[4],
		PrenomEnfant:       args[5],
		Nom:                args[6],
		Sexe:               args[7],
		De0:                args[8],
		Pre0:               args[9],
		NeLe0:              args[10],
		NeA:                args[11],
		Profession:         args[12],
		CniP:               args[13],
		De1:                args[14],
		Pre1:               args[15],
		NeLe1:              args[16],
		NeeA:               args[17],
		Profession2:        args[18],
		Dom:                args[19],
		SM:                 args[20],
		CniM:               args[21],
		Ville:              args[22],
		Tem:                args[23],
		DateDEC:            args[24],
		SignatureDeclarant: args[25],
		Signataire:         args[26],
		NumCNIDEC:          args[27],
		ProfDEC:            args[28],
		DateENREG:          time.Now(),
	}

	declarationNAsBytes, err := stub.GetState(declarationN.Numero)
	if declarationNAsBytes != nil {
		return shim.Error("declaration de naissance avec le numero suivant deja existant : " + declarationN.Numero)
	}

	versionsBytes, err := json.Marshal(declarationN)
	if err != nil {
		return shim.Error("Impossible de monter la structure de declaration de naissance : " + err.Error())
	}

	err = stub.PutState(declarationN.Numero, versionsBytes)
	if err != nil {
		return shim.Error("Impossible d'enregistrer la declaration de naissance: " + err.Error())
	}
	return shim.Success(nil)
}

func (t *Administrative_docChaincode) Recupdeclaration(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	declarationNaissanceAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Impossible de trouver cette declaration de naissance: " + err.Error())
	} else if declarationNaissanceAsBytes == nil {
		return shim.Error("Declaration de naissance: " + args[0] + " inexistante.")
	}

	return shim.Success(declarationNaissanceAsBytes)
}

func (t *Administrative_docChaincode) consulterDeclarationsNAISS(stub shim.ChaincodeStubInterface) peer.Response {
	resultiterator, err := stub.GetQueryResult(`{"selector":{"Type":"declaration de naissance"}}`)

	if err != nil {
		return shim.Error("Erreur lors de la recuperation des donées" + err.Error())
	}
	defer resultiterator.Close()

	var results []DeclarationNaissance

	for resultiterator.HasNext() {
		queryReponse, err := resultiterator.Next()
		if err != nil {
			return shim.Error("Erreur lors de la recuperation des données" + err.Error())
		}

		var fichedeclaration DeclarationNaissance
		err = json.Unmarshal(queryReponse.Value, &fichedeclaration)
		if err != nil {
			return shim.Error("Erreur lors de la deserialisation des données")
		}

		results = append(results, fichedeclaration)
	}

	resultatsAsbytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error("Erreur lors de la conversion en JSON des declarations de naissance")
	}

	return shim.Success(resultatsAsbytes)
}

// acte de deces
func (t *Administrative_docChaincode) creerActeDeces(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 24 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	acteDeces := ActeDeces{
		Numero:         args[0],
		Type:           "acte de deces",
		PROVINCE:       args[1],
		DEPARTMENT:     args[2],
		ARRONDISSEMENT: args[3],
		De:             args[4],

		NomDefunt:       args[5],
		PrenomDefunt:    args[6],
		EstDecesA:       args[7],
		DecedeLe:        args[8],
		EstNeA:          args[9],
		Le:              args[10],
		Sexe:            args[11],
		Natio:           args[12],
		Profession:      args[13],
		DomicilieDefunt: args[14],

		NP:         args[15],
		NM:         args[16],
		DomicilieA: args[17],

		DeclarationDe:     args[18],
		ParNous:           args[19],
		Centre:            args[20],
		AssisteDe:         args[21],
		SignatureOfficier: args[22],
		NumeroDEC:         args[23],
		DresseLe:          time.Now(),
	}

	actesDecesAsBytes, err := stub.GetState(acteDeces.Numero)
	if actesDecesAsBytes != nil {
		return shim.Error("Acte de decès avec le numero suivant deja existant : " + acteDeces.Numero)
	}

	Versionbyte, err := json.Marshal(acteDeces)
	if err != nil {
		return shim.Error("Impossible de monter la structure de acte de decès : " + err.Error())
	}

	err = stub.PutState(acteDeces.Numero, Versionbyte)
	if err != nil {
		return shim.Error("Impossible d'enregistrer l'acte de decès: " + err.Error())
	}
	return shim.Success(nil)
}

func (t *Administrative_docChaincode) RecupActeDeces(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	acteDecesBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Impossible de trouver cette acte de decès: " + err.Error())
	} else if acteDecesBytes == nil {
		return shim.Error("Acte de decès: " + args[0] + " inexistant.")
	}

	return shim.Success(acteDecesBytes)
}

func (t *Administrative_docChaincode) consulterActesDeces(stub shim.ChaincodeStubInterface) peer.Response {
	resultiterator, err := stub.GetQueryResult(`{"selector":{"Type":"acte de deces"}}`)

	if err != nil {
		return shim.Error("Erreur lors de la recuperation des donées" + err.Error())
	}
	defer resultiterator.Close()

	var results []ActeDeces

	for resultiterator.HasNext() {
		queryReponse, err := resultiterator.Next()
		if err != nil {
			return shim.Error("Erreur lors de la recuperation des données" + err.Error())
		}
		var ficheActe ActeDeces
		err = json.Unmarshal(queryReponse.Value, &ficheActe)
		if err != nil {
			return shim.Error("Erreur lors de la recuperation des données")
		}

		results = append(results, ficheActe)
	}

	resultatsAsbytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error("Erreur lors de la conversion en JSON des actes de decès")
	}
	return shim.Success(resultatsAsbytes)
}

// declaration de deces
func (t *Administrative_docChaincode) creerDeclarationsDeces(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 21 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	declarationD := DeclarationDeces{
		Type:              "declaration de decès",
		Nom_def:           args[0],
		Date_naissance:    args[1],
		Lieu_naissance:    args[2],
		Date_deces:        args[3],
		Heure_deces:       args[4],
		Numero:            args[5],
		Lieu_deces:        args[6],
		Cause_deces:       args[7],
		Profession:        args[8],
		Nom_temoins:       args[9],
		Ville_declaration: args[10],
		Sexe:              args[11],
		Nationalite:       args[12],
		Etat_civil:        args[13],
		Numero_identite:   args[14],
		DateELAB:          args[15],
		HeureELAB:         args[16],
		NumCNIDEC:         args[17],
		ProfDEC:           args[18],
		Signataire:        args[19],
		Signature:         args[20],
		DateENREG:         time.Now(),
	}

	declarationDAsBytes, err := stub.GetState(declarationD.Numero)
	if declarationDAsBytes != nil {
		return shim.Error("declaration de decès avec le numero suivant deja existant : " + declarationD.Numero)
	}

	versionsBytes, err := json.Marshal(declarationD)
	if err != nil {
		return shim.Error("Impossible de monter la structure de declaration de decès : " + err.Error())
	}

	err = stub.PutState(declarationD.Numero, versionsBytes)
	if err != nil {
		return shim.Error("Impossible d'enregistrer la declaration de decès: " + err.Error())
	}
	return shim.Success(nil)
}

func (t *Administrative_docChaincode) RecupDeclarationDeces(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	declarationAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Impossible de trouver cette declaration de decès: " + err.Error())
	} else if declarationAsBytes == nil {
		return shim.Error("Declaration de decès: " + args[0] + " inexistante.")
	}

	return shim.Success(declarationAsBytes)
}

func (t *Administrative_docChaincode) consulterDeclarationsDeces(stub shim.ChaincodeStubInterface) peer.Response {
	resultiterator, err := stub.GetQueryResult(`{"selector":{"Type":"declaration de decès"}}`)

	if err != nil {
		return shim.Error("Erreur lors de la recuperation des donées" + err.Error())
	}
	defer resultiterator.Close()

	var results []DeclarationDeces

	for resultiterator.HasNext() {
		queryReponse, err := resultiterator.Next()
		if err != nil {
			return shim.Error("Erreur lors de la recuperation des données" + err.Error())
		}

		var fichedeclaration DeclarationDeces
		err = json.Unmarshal(queryReponse.Value, &fichedeclaration)
		if err != nil {
			return shim.Error("Erreur lors de la deserialisation des données")
		}

		results = append(results, fichedeclaration)
	}

	resultatsAsbytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error("Erreur lors de la conversion en JSON des declarations de decès")
	}

	return shim.Success(resultatsAsbytes)
}

// publication de bans
func (t *Administrative_docChaincode) EnregistrerPUB(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 34 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	publicationB := PublicationBans{
		Numero:           args[0],
		Type:             "publication de bans",
		Date_debut:       args[1],
		Date_fin:         args[2],
		Lieu_publication: args[3],

		Nom_epoux:             args[4],
		Prenoms_epoux:         args[5],
		Date_naissance_epoux:  args[6],
		Lieu_naissance_epoux:  args[7],
		Profession_epoux:      args[8],
		Nationalite_epoux:     args[9],
		Domicile_epoux:        args[10],
		EtatCivil_epoux:       args[11],
		Numero_identite_epoux: args[12],

		Nom_epouse:             args[13],
		Prenoms_epouse:         args[14],
		Date_naissance_epouse:  args[15],
		Lieu_naissance_epouse:  args[16],
		Profession_epouse:      args[17],
		Nationalite_epouse:     args[18],
		Domicile_epouse:        args[19],
		EtatCivil_epouse:       args[20],
		Numero_identite_epouse: args[21],

		Date_mariage:       args[22],
		HeureDebut_mariage: args[23],
		DetailsMariage:     args[24],
		NomPereMarieHomme:  args[25],

		NomMereMarieHomme: args[26],
		NomPereMarieFemme: args[27],
		NomMereMarieFemme: args[28],
		MunicipalitePub:   args[29],

		Signataire:        args[30],
		Signature:         args[31],
		DateELAB:          args[32],
		ContactMunicipale: args[33],

		DateENREG: time.Now(),
	}

	publicationBAsBytes, err := stub.GetState(publicationB.Numero)
	if publicationBAsBytes != nil {
		return shim.Error("declaration de naissance avec le numero suivant deja existant : " + publicationB.Numero)
	}

	versionsBytes, err := json.Marshal(publicationB)
	if err != nil {
		return shim.Error("Impossible de monter la structure de la publication de bans : " + err.Error())
	}

	err = stub.PutState(publicationB.Numero, versionsBytes)
	if err != nil {
		return shim.Error("Impossible d'enregistrer la publication de bans: " + err.Error())
	}
	return shim.Success(nil)
}

func (t *Administrative_docChaincode) RecupererPublicationBans(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	publicationAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Impossible de trouver cette publication de bans: " + err.Error())
	} else if publicationAsBytes == nil {
		return shim.Error("Publication de bans: " + args[0] + " inexistante.")
	}

	return shim.Success(publicationAsBytes)
}

func (t *Administrative_docChaincode) ConsulterPubBans(stub shim.ChaincodeStubInterface) peer.Response {
	resultiterator, err := stub.GetQueryResult(`{"selector":{"Type":"publication de bans"}}`)

	if err != nil {
		return shim.Error("Erreur lors de la recuperation des données" + err.Error())
	}
	defer resultiterator.Close()

	var results []PublicationBans

	for resultiterator.HasNext() {
		queryReponse, err := resultiterator.Next()
		if err != nil {
			return shim.Error("Erreur lors de la recuperation des données" + err.Error())
		}

		var fichepublication PublicationBans
		err = json.Unmarshal(queryReponse.Value, &fichepublication)
		if err != nil {
			return shim.Error("Erreur lors de la deserialisation des données")
		}

		results = append(results, fichepublication)
	}

	resultatsAsbytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error("Erreur lors de la conversion en JSON des publications de bans")
	}

	return shim.Success(resultatsAsbytes)
}

// acte de mariage
func (t *Administrative_docChaincode) CreerActesMariage(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 46 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	acteMariage := ActeMariage{
		Numero:         args[0],
		Type:           "acte de mariage",
		PROVINCE:       args[1],
		DEPARTMENT:     args[2],
		ARRONDISSEMENT: args[3],
		De:             args[4],

		Contracter: args[5],

		NomEpoux:     args[6],
		PrenomEpoux:  args[7],
		NomPereH:     args[8],
		NomMereH:     args[9],
		DateNEpoux:   args[10],
		LieuNEpoux:   args[11],
		RaceH:        args[12],
		GroupemmentH: args[13],
		SubdivisionH: args[14],
		RegionH:      args[15],
		ProffEPH:     args[16],
		ResidenceH:   args[17],

		NomEpouse:    args[18],
		PrenomEpouse: args[19],
		NomPereF:     args[20],
		NomMereF:     args[21],
		DateNEpouse:  args[22],
		LieuNEpouse:  args[23],
		RaceF:        args[24],
		GroupemmentF: args[25],
		SubdivisionF: args[26],
		RegionF:      args[27],
		ProffEPF:     args[28],
		ResidenceF:   args[29],

		ConsentementEPOUX:   args[30],
		ConsentementCHEFFam: args[31],
		ConsentementEpoux:   args[32],
		ConsentementEpouse:  args[33],
		Oppositions:         args[34],
		MontantDOT:          args[35],
		SommeDOTVERSER:      args[36],
		DateVerse:           args[37],
		DateVerseCOM:        args[38],
		TemoinsH:            args[39],
		TemoinsF:            args[40],

		ParNous:           args[41],
		Centre:            args[42],
		AssisteDe:         args[43],
		SignatureOfficier: args[44],
		NumeroDEC:         args[45],

		DresseLe: time.Now(),
	}

	actesMariageAsBytes, err := stub.GetState(acteMariage.Numero)
	if actesMariageAsBytes != nil {
		return shim.Error("Acte de mariage avec le numero suivant deja existant : " + acteMariage.Numero)
	}

	Versionbyte, err := json.Marshal(acteMariage)
	if err != nil {
		return shim.Error("Impossible de monter la structure de l'acte de mariage : " + err.Error())
	}

	err = stub.PutState(acteMariage.Numero, Versionbyte)
	if err != nil {
		return shim.Error("Impossible d'enregistrer l'acte de mariage: " + err.Error())
	}
	return shim.Success(nil)
}

func (t *Administrative_docChaincode) RecupActeMariage(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Nombre d'arguments insuffisant.Veuiller reessayer avec un nombre d'argument coherent.")
	}

	acteMariageAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Impossible de trouver cette acte de mariage: " + err.Error())
	} else if acteMariageAsBytes == nil {
		return shim.Error("Acte de mariage: " + args[0] + " inexistant.")
	}

	return shim.Success(acteMariageAsBytes)
}

func (t *Administrative_docChaincode) consulterActesMariage(stub shim.ChaincodeStubInterface) peer.Response {
	resultiterator, err := stub.GetQueryResult(`{"selector":{"Type":"acte de mariage"}}`)

	if err != nil {
		return shim.Error("Erreur lors de la recuperation des donées" + err.Error())
	}
	defer resultiterator.Close()

	var results []ActeMariage

	for resultiterator.HasNext() {
		queryReponse, err := resultiterator.Next()
		if err != nil {
			return shim.Error("Erreur lors de la recuperation des données" + err.Error())
		}
		var ficheActe ActeMariage
		err = json.Unmarshal(queryReponse.Value, &ficheActe)
		if err != nil {
			return shim.Error("Erreur lors de la recuperation des données")
		}

		results = append(results, ficheActe)
	}

	resultatsAsbytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error("Erreur lors de la conversion en JSON des actes de mariage")
	}
	return shim.Success(resultatsAsbytes)
}

// function principale
func main() {
	fmt.Printf("Chargement du contrat intelligent administrative_doc...")
	err := shim.Start(new(Administrative_docChaincode))
	if err != nil {
		fmt.Printf("Impossible de charger le contrat intelligent administrative_doc : %s", err)
	}
}
