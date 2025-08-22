#!/bin/bash
# Safe i18n string updates for CyFir

echo "=== Updating i18n Translation Files ===="
echo "Only updating user-visible brand names, not technical terms"
echo

# French translations
echo "Updating French translations..."
sed -i 's/"Velociraptor Binary":"Binaire Velociraptor"/"Velociraptor Binary":"Binaire CyFir"/g' gui/velociraptor/src/components/i8n/fr.jsx
sed -i 's/mise à jour du serveur Velociraptor/mise à jour du serveur CyFir/g' gui/velociraptor/src/components/i8n/fr.jsx
sed -i 's/"Default Velociraptor":"Vélociraptor standard"/"Default Velociraptor":"CyFir standard"/g' gui/velociraptor/src/components/i8n/fr.jsx
sed -i 's/"Velociraptor Classic (light)": "Vélociraptor Classique (léger)"/"Velociraptor Classic (light)": "CyFir Classique (léger)"/g' gui/velociraptor/src/components/i8n/fr.jsx
sed -i 's/"Velociraptor (light)":"Vélociraptor (léger)"/"Velociraptor (light)":"CyFir (léger)"/g' gui/velociraptor/src/components/i8n/fr.jsx
sed -i 's/"Velociraptor (dark)":"Vélociraptor (foncé)"/"Velociraptor (dark)":"CyFir (foncé)"/g' gui/velociraptor/src/components/i8n/fr.jsx
sed -i 's/serveurs Velociraptor/serveurs CyFir/g' gui/velociraptor/src/components/i8n/fr.jsx
sed -i 's/car Velociraptor est/car CyFir est/g' gui/velociraptor/src/components/i8n/fr.jsx
echo "French translations updated"

# German translations
echo
echo "Updating German translations..."
sed -i 's/"Velociraptor Binary":"Velociraptor Binärdatei"/"Velociraptor Binary":"CyFir Binärdatei"/g' gui/velociraptor/src/components/i8n/de.jsx
sed -i 's/"Default Velociraptor":"Standard Velociraptor"/"Default Velociraptor":"Standard CyFir"/g' gui/velociraptor/src/components/i8n/de.jsx
sed -i 's/"Velociraptor Classic (light)": "Velociraptor Klassisch (hell)"/"Velociraptor Classic (light)": "CyFir Klassisch (hell)"/g' gui/velociraptor/src/components/i8n/de.jsx
sed -i 's/"Velociraptor (light)":"Velociraptor (hell)"/"Velociraptor (light)":"CyFir (hell)"/g' gui/velociraptor/src/components/i8n/de.jsx
sed -i 's/"Velociraptor (dark)":"Velociraptor (dunkel)"/"Velociraptor (dark)":"CyFir (dunkel)"/g' gui/velociraptor/src/components/i8n/de.jsx
sed -i 's/Velociraptor-Server/CyFir-Server/g' gui/velociraptor/src/components/i8n/de.jsx
echo "German translations updated"

# Spanish translations
echo
echo "Updating Spanish translations..."
sed -i 's/"Velociraptor Binary":"Binario Velociraptor"/"Velociraptor Binary":"Binario CyFir"/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/"Default Velociraptor":"Velociraptor estándar"/"Default Velociraptor":"CyFir estándar"/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/"Velociraptor Classic (light)": "Velociraptor Clásico (claro)"/"Velociraptor Classic (light)": "CyFir Clásico (claro)"/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/"Velociraptor (light)":"Velociraptor (claro)"/"Velociraptor (light)":"CyFir (claro)"/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/"Velociraptor (dark)":"Velociraptor (oscuro)"/"Velociraptor (dark)":"CyFir (oscuro)"/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/servidor de Velociraptor/servidor de CyFir/g' gui/velociraptor/src/components/i8n/es.jsx
echo "Spanish translations updated"

# Portuguese translations
echo
echo "Updating Portuguese translations..."
sed -i 's/"Velociraptor Binary":"Binário Velociraptor"/"Velociraptor Binary":"Binário CyFir"/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/"Default Velociraptor":"Velociraptor padrão"/"Default Velociraptor":"CyFir padrão"/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/"Velociraptor Classic (light)": "Velociraptor Clássico (claro)"/"Velociraptor Classic (light)": "CyFir Clássico (claro)"/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/"Velociraptor (light)":"Velociraptor (claro)"/"Velociraptor (light)":"CyFir (claro)"/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/"Velociraptor (dark)":"Velociraptor (escuro)"/"Velociraptor (dark)":"CyFir (escuro)"/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/servidor Velociraptor/servidor CyFir/g' gui/velociraptor/src/components/i8n/por.jsx
echo "Portuguese translations updated"

# Japanese translations
echo
echo "Updating Japanese translations..."
sed -i 's/"Velociraptor Binary":"Velociraptorバイナリ"/"Velociraptor Binary":"CyFirバイナリ"/g' gui/velociraptor/src/components/i8n/jp.jsx
sed -i 's/"Default Velociraptor":"デフォルトVelociraptor"/"Default Velociraptor":"デフォルトCyFir"/g' gui/velociraptor/src/components/i8n/jp.jsx
sed -i 's/"Velociraptor Classic (light)": "Velociraptorクラシック（ライト）"/"Velociraptor Classic (light)": "CyFirクラシック（ライト）"/g' gui/velociraptor/src/components/i8n/jp.jsx
sed -i 's/"Velociraptor (light)":"Velociraptor（ライト）"/"Velociraptor (light)":"CyFir（ライト）"/g' gui/velociraptor/src/components/i8n/jp.jsx
sed -i 's/"Velociraptor (dark)":"Velociraptor（ダーク）"/"Velociraptor (dark)":"CyFir（ダーク）"/g' gui/velociraptor/src/components/i8n/jp.jsx
sed -i 's/Velociraptorサーバー/CyFirサーバー/g' gui/velociraptor/src/components/i8n/jp.jsx
echo "Japanese translations updated"

# Vietnamese translations
echo
echo "Updating Vietnamese translations..."
sed -i 's/"Velociraptor Binary":"Tệp nhị phân Velociraptor"/"Velociraptor Binary":"Tệp nhị phân CyFir"/g' gui/velociraptor/src/components/i8n/vi.jsx
sed -i 's/"Default Velociraptor":"Velociraptor mặc định"/"Default Velociraptor":"CyFir mặc định"/g' gui/velociraptor/src/components/i8n/vi.jsx
sed -i 's/"Velociraptor Classic (light)": "Velociraptor Cổ điển (sáng)"/"Velociraptor Classic (light)": "CyFir Cổ điển (sáng)"/g' gui/velociraptor/src/components/i8n/vi.jsx
sed -i 's/"Velociraptor (light)":"Velociraptor (sáng)"/"Velociraptor (light)":"CyFir (sáng)"/g' gui/velociraptor/src/components/i8n/vi.jsx
sed -i 's/"Velociraptor (dark)":"Velociraptor (tối)"/"Velociraptor (dark)":"CyFir (tối)"/g' gui/velociraptor/src/components/i8n/vi.jsx
sed -i 's/máy chủ Velociraptor/máy chủ CyFir/g' gui/velociraptor/src/components/i8n/vi.jsx
echo "Vietnamese translations updated"

echo
echo "=== i18n Updates Complete ==="
echo "All translation files have been updated with CyFir branding"