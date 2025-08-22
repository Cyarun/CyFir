#\!/bin/bash
# Fix remaining i18n translations

echo "=== Fixing Remaining i18n Translations ==="

# German
echo "Fixing German translations..."
sed -i 's/"Velociraptor Binary":"Velociraptor Binary"/"Velociraptor Binary":"CyFir Binary"/g' gui/velociraptor/src/components/i8n/de.jsx
sed -i 's/lädt Velociraptor es von/lädt CyFir es von/g' gui/velociraptor/src/components/i8n/de.jsx
sed -i 's/da Velociraptor nicht/da CyFir nicht/g' gui/velociraptor/src/components/i8n/de.jsx
sed -i 's/"Default Velociraptor":"Standard-Velociraptor"/"Default Velociraptor":"Standard-CyFir"/g' gui/velociraptor/src/components/i8n/de.jsx
sed -i 's/"Velociraptor Classic (light)": "Velociraptor-Klassiker (hell)"/"Velociraptor Classic (light)": "CyFir-Klassiker (hell)"/g' gui/velociraptor/src/components/i8n/de.jsx

# Spanish
echo "Fixing Spanish translations..."
grep -n "Velociraptor" gui/velociraptor/src/components/i8n/es.jsx  < /dev/null |  grep -v ":" | head -5
sed -i 's/"Velociraptor Binary":"Velociraptor Binary"/"Velociraptor Binary":"CyFir Binary"/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/Velociraptor lo descargará/CyFir lo descargará/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/ya que Velociraptor no puede/ya que CyFir no puede/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/"Default Velociraptor":"Velociraptor predeterminado"/"Default Velociraptor":"CyFir predeterminado"/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/"Velociraptor Classic (light)": "Velociraptor Clásico (claro)"/"Velociraptor Classic (light)": "CyFir Clásico (claro)"/g' gui/velociraptor/src/components/i8n/es.jsx

# Portuguese
echo "Fixing Portuguese translations..."
sed -i 's/"Velociraptor Binary":"Velociraptor Binary"/"Velociraptor Binary":"CyFir Binary"/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/o Velociraptor baixará/o CyFir baixará/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/pois o Velociraptor não consegue/pois o CyFir não consegue/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/"Default Velociraptor":"Velociraptor padrão"/"Default Velociraptor":"CyFir padrão"/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/"Velociraptor Classic (light)": "Velociraptor Clássico (claro)"/"Velociraptor Classic (light)": "CyFir Clássico (claro)"/g' gui/velociraptor/src/components/i8n/por.jsx

# Japanese
echo "Fixing Japanese translations..."
sed -i 's/"Velociraptor Binary":"Velociraptor Binary"/"Velociraptor Binary":"CyFir Binary"/g' gui/velociraptor/src/components/i8n/jp.jsx
sed -i 's/Velociraptorは/CyFirは/g' gui/velociraptor/src/components/i8n/jp.jsx
sed -i 's/"Default Velociraptor":"デフォルトのVelociraptor"/"Default Velociraptor":"デフォルトのCyFir"/g' gui/velociraptor/src/components/i8n/jp.jsx
sed -i 's/"Velociraptor Classic (light)": "Velociraptor クラシック（ライト）"/"Velociraptor Classic (light)": "CyFir クラシック（ライト）"/g' gui/velociraptor/src/components/i8n/jp.jsx

# Vietnamese
echo "Fixing Vietnamese translations..."
sed -i 's/"Velociraptor Binary":"Velociraptor Binary"/"Velociraptor Binary":"CyFir Binary"/g' gui/velociraptor/src/components/i8n/vi.jsx
sed -i 's/Velociraptor sẽ tải/CyFir sẽ tải/g' gui/velociraptor/src/components/i8n/vi.jsx
sed -i 's/vì Velociraptor không thể/vì CyFir không thể/g' gui/velociraptor/src/components/i8n/vi.jsx

# French (final check)
echo "Fixing French translations..."
sed -i 's/"Velociraptor Binary":"Velociraptor Binary"/"Velociraptor Binary":"CyFir Binary"/g' gui/velociraptor/src/components/i8n/fr.jsx

echo "=== Remaining i18n Fixes Complete ==="
